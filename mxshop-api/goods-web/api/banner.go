package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"goods-web/forms"
	"goods-web/global"
	"goods-web/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"net/http"
	"strconv"
)

func ListBanner(ctx *gin.Context) {
	rsp, err := global.GoodsSrvClient.BannerList(context.Background(), &emptypb.Empty{})
	if err != nil {
		HandleGrpcErrorToHttp(err, ctx)
		return
	}
	result := make([]interface{}, 0)
	for _, value := range rsp.Data {
		reMap := make(map[string]interface{})
		reMap["id"] = value.Id
		reMap["index"] = value.Index
		reMap["image"] = value.Image
		reMap["url"] = value.Url
		result = append(result, reMap)
	}
	ctx.JSON(http.StatusOK, result)
}

func NewBanner(ctx *gin.Context) {
	bannerForm := forms.BannerForm{}
	if err := ctx.ShouldBindJSON(&bannerForm); err != nil {
		HandleValidatorError(ctx, err)
		return
	}

	rsp, err := global.GoodsSrvClient.CreateBanner(context.Background(),
		&proto.BannerRequest{
			Index: int32(bannerForm.Index),
			Url:   bannerForm.Url,
			Image: bannerForm.Image,
		})

	if err != nil {
		HandleGrpcErrorToHttp(err, ctx)
		return
	}
	response := make(map[string]interface{})
	response["id"] = rsp.Id
	response["index"] = rsp.Index
	response["url"] = rsp.Url
	response["image"] = rsp.Image

	ctx.JSON(http.StatusOK, response)
}

func UpdateBanner(ctx *gin.Context) {
	bannerForm := forms.BannerForm{}
	if err := ctx.ShouldBindJSON(&bannerForm); err != nil {
		HandleValidatorError(ctx, err)
		return
	}
	id := ctx.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		ctx.Status(http.StatusNotFound)
		return
	}
	_, err = global.GoodsSrvClient.UpdateBanner(context.Background(), &proto.BannerRequest{
		Id:    int32(idInt),
		Index: int32(bannerForm.Index),
		Url:   bannerForm.Url,
	})
	if err != nil {
		HandleGrpcErrorToHttp(err, ctx)
		return
	}
	ctx.Status(http.StatusOK)
}

func DeleteBanner(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		ctx.Status(http.StatusNotFound)
		return
	}
	_, err = global.GoodsSrvClient.DeleteBanner(context.Background(), &proto.BannerRequest{Id: int32(idInt)})
	if err != nil {
		HandleGrpcErrorToHttp(err, ctx)
		return
	}
	ctx.Status(http.StatusOK)
}
