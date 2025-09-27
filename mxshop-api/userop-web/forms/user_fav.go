package forms

type UserFavForm struct {
	GoodsId uint64 `form:"goodsId" json:"goodsId" binding:"required"`
}
