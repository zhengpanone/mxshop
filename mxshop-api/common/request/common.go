package request

type PageRequest struct {
	Page     int `json:"page" form:"page" binding:"required,min=1"`                 // 页码
	PageSize int `json:"pageSize" form:"pageSize" binding:"required,min=1,max=100"` // 每页大小
}
type CommonIds struct {
	Ids []interface{} `json:"ids"`
}

type Empty struct{}
