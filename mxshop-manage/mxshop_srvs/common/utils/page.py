from typing import Union, Optional

from common.proto.pb.common_pb2 import PageResponse,PageRequest


def make_page_response(total:int,page:Union[PageRequest,int], page_size:Optional[int] = None)->PageResponse:
    from math import ceil
    rsp = PageResponse()
    rsp.total = total

    if isinstance(page,PageRequest):
        rsp.pageNum = page.pageNum
        rsp.pageSize = page.pageSize
    else:
        if page_size is None:
            raise ValueError("page_size must be provided when page is int")
        rsp.pageNum = page
        rsp.pageSize = page_size
    rsp.totalPage = ceil(total/rsp.pageSize)

    return rsp
