package forms

type ShopCartItemForm struct {
	GoodsId uint64 `form:"goods" json:"goods" binding:"required"`
	Nums    uint64 `form:"nums" json:"nums" binding:"required,min=1"`
}

type ShopCartItemUpdateForm struct {
	Nums    uint64 `form:"nums" json:"nums" binding:"required,min=1"`
	Checked *bool  `form:"checked" json:"checked"`
}
