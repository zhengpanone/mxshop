/**
 * 公共基础接口封装
 */
import request from '@/utils/request'
import { IResponseData, ILoginInfo, ILoginResponse, ILoginForm, ILoginPayload, ICaptcha } from './types/common'

let userUrl = "http://127.0.0.1:18021"

export const getLoginInfo = () => {
  return request<IResponseData<ILoginInfo>>({
    method: 'GET',
    url: '/api/user/login/info',
  })
}
/**
 * 获取图片验证码
 * @returns 返回图片验证码
 */
export const getCaptcha = () => {
  return request<IResponseData<ICaptcha>>({
    url: userUrl + '/v1/base/captcha',
    method: 'get'
  })
}

export const login = (data: ILoginPayload) => {
  return request<IResponseData<ILoginResponse>>({
    url: userUrl + '/v1/user/pwd_login',
    method: 'POST',
    data: data
  })
}


export const ossPolicy = () => {
  return request({
    url: '/oss/token',
    method: 'get',
  })
}
