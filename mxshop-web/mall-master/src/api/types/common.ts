export interface IResponseData<T> {
  code: number
  msg: string
  data: T
}

// 定义完整的 API 响应结构
export interface IResponsePageData<T> {
  code: number;
  msg: string;
  data: {
    list: T[];
    pageNum: number;
    pageSize: number;
    total: number;
    totalPage: number;
  }
}


// 定义分页和总条数
export interface IPagination {
  pageNum: number;
  pageSize: number;
  total: number;
  totalPage: number;
}



/**
 * 验证码接口返回数据
 */
export interface ICaptcha {
  captchaId: string;
  imageBase64: string;
}

/**
 * 登录接口参数
 */
export interface ILoginForm {
  account: string
  password: string
  captchaText: string
}

/**
 * 登录请求参数
 */
export interface ILoginPayload extends ILoginForm {
  captchaId: string
}

export interface ILoginInfo {
  slide: string[]
  logoSquare: string
  logoRectangle: string
  loginLogo: string
}

export interface IUserInfo {
  id: number
  account: string
  head_pic: string
}

export interface IMenu {
  path: string
  title: string
  icon: string
  header: string
  is_header: number
  children?: IMenu[]
}

export interface ILoginResponse {
  token: string
  expires_time: number
  menus: IMenu[]
  unique_auth: string[]
  userInfo: IUserInfo
  logo: string
  logo_square: string
  version: string
  newOrderAudioLink: string
}

export interface IDict {
  code: string
  name: string
}
