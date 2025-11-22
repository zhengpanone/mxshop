package controller

import "github.com/gin-gonic/gin"

// MenuController 菜单管理
type MenuController struct{}

func (m *MenuController) GetMenuList(ctx *gin.Context) {}

func (m *MenuController) GetMenuTree(ctx *gin.Context) {}

func (m *MenuController) GetMenuTreeById(ctx *gin.Context) {}

func (m *MenuController) AddMenu(ctx *gin.Context) {

}

func (m *MenuController) UpdateMenu(ctx *gin.Context) {}

func (m *MenuController) DeleteMenu(ctx *gin.Context) {}

func (m *MenuController) GetMenuListByRoleId(ctx *gin.Context) {}
