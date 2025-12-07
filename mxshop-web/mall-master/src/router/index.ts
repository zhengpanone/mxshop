import { createRouter, createWebHashHistory, RouteRecordRaw } from 'vue-router'
import GoodsRoutes from './modules/goods'
import OrderRoutes from './modules/order'
import UmsRoutes from './modules/ums'
import AppLayout from '@/views/layout/AppLayout.vue'
import nprogress from 'nprogress'
import 'nprogress/nprogress.css'
import { indexStore } from '@/store/index'

// import eventEmiter from '@/utils/eventEmiter'

const routes: RouteRecordRaw[] = [
  {
    path: '/',
    component: AppLayout,
    meta: {
      requireAuth: true,
    },
    children: [
      {
        path: '', // 默认子路由
        name: 'home',
        component: () => import('@/views/home/index.vue'),
        meta: {
          title: '首页',
        },
      },
      GoodsRoutes,
      OrderRoutes,
      UmsRoutes,
    ],
  },
  {
    path: '/login',
    name: 'login',
    component: () => import('@/views/login/index.vue'),
    meta: {
      title: '登录'
    }
  },
  {
    path: '/forget_password',
    name: 'forgetPassword',
    component: () => import('@/views/forgetPassword/index.vue'),
    meta: {
      title: '忘记密码'
    }
  },
] // 路由规则

const router = createRouter({
  history: createWebHashHistory(), // 路由模式
  routes,
})
// 全局前置守卫
router.beforeEach((to, from) => {
  const store = indexStore()
  if (to.meta.requireAuth && !store.$state.user) {
    return {
      path: '/login',
      query: { redirect: to.fullPath },
    }
  }
  nprogress.start()
})

router.afterEach(() => {
  nprogress.done()
})

// eventEmiter.on('API:UN_AUTH',()=>{
//   router.push('/')
// })

export default router
