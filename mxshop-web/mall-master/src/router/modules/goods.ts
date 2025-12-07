import { RouteRecordRaw, RouterView } from 'vue-router'

const routes: RouteRecordRaw = {
  path: '/pms',
  name: 'pms',
  component: RouterView,
  meta: {
    title: '商品',
    icon: 'goods'
  },
  children: [
    {
      path: 'goods_list',
      name: 'goodsList',
      component: () => import('@/views/pms/goods/index.vue'),
      meta: {
        // 自定义元数据
        title: '商品列表',
        icon: 'goods-list'
      },
    },
    {
      path: 'product_add',
      name: 'productAdd',
      component: () => import('@/views/pms/goods/add.vue'),
      meta: { title: '添加商品', icon: 'product-add', hidden: true },


    },
    {
      path: 'product_category',
      name: 'productCategory',
      component: () => import('@/views/pms/productCategory/index.vue'),
      meta: {
        title: '商品分类',
      },
    },
    {
      path: 'product_category_add',
      name: 'productCategoryAdd',
      component: () => import('@/views/pms/productCategory/add.vue'),
      meta: { title: '添加商品分类', hidden: true },
    },
    {
      path: 'product_brand',
      name: 'productBrand',
      component: () => import('@/views/pms/productBrand/index.vue'),
      meta: {
        title: '商品品牌',
      },
    },
    {
      path: 'product_brand_add',
      name: 'productBrandAdd',
      component: () => import('@/views/pms/productBrand/add.vue'),
      meta: { title: '添加商品品牌', hidden: true },

    },
    // {
    //   path: 'updateBrand',
    //   name: 'updateBrand',
    //   component: () => import('@/views/pms/brand/update'),
    //   meta: {title: '编辑品牌'hidden: true},

    // },
    {
      path: 'product_attr',
      name: 'product_attr',
      component: () => import('@/views/pms/productAttr/index.vue'),
      meta: {
        title: '商品类型',
      },
    },
    /*{
      path: 'product_reply',
      name: 'product_reply',
      component: () => import('@/views/pms/reply/index.vue'),
      meta: {
        title: '商品评论',
      },
    },*/
  ],
}

export default routes
