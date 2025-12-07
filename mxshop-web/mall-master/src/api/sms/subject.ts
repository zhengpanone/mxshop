import request from '@/utils/request'
export const fetchSubjectListAll = () => {
  return request({
    url: '/subject/listAll',
    method: 'get',
  })
}

export const fetchSubjectList = (params: object) => {
  return request({
    url: '/subject/list',
    method: 'get',
    params: params
  })
}
