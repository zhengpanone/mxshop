package utils

import (
	commonpb "github.com/zhengpanone/mxshop/mxshop-api/common/proto/pb"
	"github.com/zhengpanone/mxshop/mxshop-api/common/response"
)

func ConvertPage[T any](page *commonpb.PageResponse, list []T) response.PageResult[T] {
	return response.PageResult[T]{
		List:     list,
		Total:    page.Total,
		PageNum:  page.PageNum,
		PageSize: page.PageSize,
	}
}
