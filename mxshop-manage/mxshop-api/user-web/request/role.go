package request

type AddRoleForm struct {
	Name   string `form:"name" json:"name" binding:"required,min=3,max=20"`
	Desc   string `form:"desc" json:"desc" `
	Status bool   `form:"status" json:"status"`
}
