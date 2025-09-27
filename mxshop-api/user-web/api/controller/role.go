package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	commonpb "github.com/zhengpanone/mxshop/mxshop-api/common/proto/pb"
	commonResponse "github.com/zhengpanone/mxshop/mxshop-api/common/response"
	commonUtils "github.com/zhengpanone/mxshop/mxshop-api/common/utils"
	"github.com/zhengpanone/mxshop/mxshop-api/user-web/global"
	"github.com/zhengpanone/mxshop/mxshop-api/user-web/request"
	"go.uber.org/zap"
)

func CreateRole(ctx *gin.Context) {
	addRoleForm := request.AddRoleForm{}
	if err := ctx.ShouldBindJSON(&addRoleForm); err != nil {
		HandleValidatorError(ctx, err)
		return
	}
	role, err := global.RoleSrvClient.CreateRole(context.Background(), &commonpb.CreateRoleRequest{
		Name:   addRoleForm.Name,
		Remark: addRoleForm.Desc,
		Status: addRoleForm.Status,
	})
	if err != nil {
		zap.S().Errorf("[CreateRole]新建角色失败：%s", err.Error())
		commonUtils.HandleGrpcErrorToHttp(err, ctx, "用户服务srv")
		return
	}
	commonResponse.OkWithData(ctx, role)
}
