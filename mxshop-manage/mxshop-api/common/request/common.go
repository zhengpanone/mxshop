package request

type PageRequest struct {
	PageNum  uint64 `json:"pageNum" form:"pageNum" binding:"required,min=1"`           // 页码
	PageSize uint64 `json:"pageSize" form:"pageSize" binding:"required,min=1,max=100"` // 每页大小
}
type CommonIds struct {
	Ids []uint64 `json:"ids" form:"ids"`
}

type Empty struct{}
