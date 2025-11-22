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

// DictTypeController 字典类型管理
type DictTypeController struct{}

func (m *MenuController) GetDictTypeList(ctx *gin.Context) {}

func (m *MenuController) GetDictTypeTree(ctx *gin.Context) {}

func (m *MenuController) GetDictTypeById(ctx *gin.Context) {}

func (d *DictTypeController) CreateDictType(ctx *gin.Context) {
	addDictTypeForm := request.CreateDictTypeRequest{}
	if err := ctx.ShouldBindJSON(&addDictTypeForm); err != nil {
		commonUtils.HandleValidatorError(ctx, global.Trans, err)
		return
	}
	dictType, err := global.DictSrvClient.CreateDictType(context.Background(), &commonpb.CreateDictTypeRequest{
		SystemFlag: addDictTypeForm.SystemFlag,
		DictCode:   addDictTypeForm.DictCode,
		DictName:   addDictTypeForm.DictName,
		Remark:     addDictTypeForm.Remark,
		Status:     addDictTypeForm.Status,
	})
	if err != nil {
		zap.S().Errorf("[CreateDictType]新建字典类型失败：%s", err.Error())
		commonUtils.HandleGrpcErrorToHttp(err, ctx, "用户服务srv")
		return
	}
	commonResponse.OkWithData(ctx, dictType)
}

func (m *MenuController) UpdateDictType(ctx *gin.Context) {}

func (m *MenuController) DeleteDictType(ctx *gin.Context) {}
