package forms

type UserFavForm struct {
	GoodsId int32 `form:"goodsId" json:"goodsId" binding:"required"`
}
