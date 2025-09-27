package response

type CategoryResponse struct {
	Name string `form:"name" json:"name" `
	Id   uint64 `json:"id" form:"id" `
}

type BrandResponse struct {
	Id   uint64 `json:"id" form:"id" `
	Name string `form:"name" json:"name" ` // 品牌名称
	Logo string `form:"logo" json:"logo"`  // 品牌logo
}
type GoodsResponse struct {
	Id          uint64           `json:"id" form:"id" binding:"required"`
	Name        string           `form:"name" json:"name" binding:"required,min=2,max=100"`
	Description string           `form:"description" json:"description" binding:"required,min=2,max=100"`
	GoodsSn     string           `form:"goods_sn" json:"goods_sn" binding:"required,min=2,lt=20"`
	Stocks      int32            `form:"stocks" json:"stocks" binding:"required,min=1"`
	Category    CategoryResponse `form:"category" json:"category" `
	MarketPrice float32          `form:"market_price" json:"market_price" binding:"required,min=0"`
	ShopPrice   float32          `form:"shop_price" json:"shop_price" binding:"required,min=0"`
	GoodsBrief  string           `form:"goods_brief" json:"goods_brief" binding:"required,min=3"`
	GoodsDesc   string           `form:"goods_desc" json:"goods_desc" binding:"required,min=3"`
	Images      []string         `form:"images" json:"images" binding:"required,min=1"`
	DescImages  []string         `form:"desc_images" json:"desc_images" binding:"required,min=1"`
	ShipFree    bool             `form:"ship_free" json:"ship_free" binding:"required"`
	FrontImage  string           `form:"front_image" json:"front_image" binding:"required,url"`
	Brand       BrandResponse    `form:"brand" json:"brand" binding:"required"`
	IsHot       bool             `form:"is_hot" json:"is_hot" binding:"required"`
	IsNew       bool             `form:"is_new" json:"is_new" binding:"required"`
	OnSale      bool             `form:"on_sale" json:"on_sale" binding:"required"`
}
