import Cookies from "js-cookie";
import type { Router, RouteLocationNormalizedLoaded, LocationQueryValue } from "vue-router";

const SupportKey = 'supportKey' as const;

export const getSupport = () => {
  return Cookies.get(SupportKey)
}

export const setSupport = (isSupport: any) => {
  return Cookies.set(SupportKey, isSupport, { expires: 3 })
}

export const setCookie = (key: any, value: any, expires: any) => {
  return Cookies.set(key, value, { expires: expires })
}

export const getCookie = (key: any) => {
  return Cookies.get(key)
}



/**
 * 通用跳转
 * @param {object} route   - 当前 route 对象
 * @param {object} router  - router 实例
 * @param {string|undefined} target - 目标路由（可选）
 * @param {string} fallback - 兜底路由（默认 "/")
 * @param {boolean} replace - 是否使用 replace 跳转（默认 false）
 */
export const navigateTo = (route: RouteLocationNormalizedLoaded,
  router: Router,
  target?: string,
  fallback: string = "/",
  replace: boolean = false): void => {
  // redirect 可以是 string | LocationQueryValue[] | undefined
  let redirect: string | LocationQueryValue | LocationQueryValue[] | undefined =
    target || route.query.redirect || fallback;


  // 如果 redirect 是数组，取第一个
  if (Array.isArray(redirect)) {
    redirect = redirect[0];
  }

  // 必须是字符串并且以 / 开头，否则用 fallback
  if (typeof redirect !== "string" || !redirect.startsWith("/")) {
    redirect = fallback;
  }

  if (replace) {
    router.replace(redirect);
  } else {
    router.push(redirect);
  }
};