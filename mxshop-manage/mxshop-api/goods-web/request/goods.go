package request

import commonRequest "github.com/zhengpanone/mxshop/mxshop-api/common/request"

type GoodsForm struct {
	Name        string   `form:"name" json:"name" binding:"required,min=2,max=100"`
	GoodsSn     string   `form:"goods_sn" json:"goods_sn" binding:"required,min=2,lt=20"`
	Stocks      int32    `form:"stocks" json:"stocks" binding:"required,min=1"`
	CategoryId  uint64   `form:"category" json:"category" binding:"required"`
	MarketPrice float32  `form:"market_price" json:"market_price" binding:"required,min=0"`
	ShopPrice   float32  `form:"shop_price" json:"shop_price" binding:"required,min=0"`
	GoodsBrief  string   `form:"goods_brief" json:"goods_brief" binding:"required,min=3"`
	GoodsDesc   string   `form:"goods_desc" json:"goods_desc" binding:"required,min=3"`
	Images      []string `form:"images" json:"images" binding:"required,min=1"`
	DescImages  []string `form:"desc_images" json:"desc_images" binding:"required,min=1"`
	ShipFree    *bool    `form:"ship_free" json:"ship_free" binding:"required"`
	FrontImage  string   `form:"front_image" json:"front_image" binding:"required,url"`
	BrandId     uint64   `form:"brand" json:"brand" binding:"required"`
}

type GoodsStatusForm struct {
	IsNew  *bool `form:"new" json:"new" binding:"required"`
	IsHot  *bool `form:"hot" json:"hot" binding:"required"`
	OnSale *bool `form:"sale" json:"sale" binding:"required"`
}

type GoodsPageForm struct {
	PriceMin                  float32 `form:"price_min" json:"price_min" binding:"min=0"`
	PriceMax                  float32 `form:"price_max" json:"price_max" binding:"min=0"`
	IsHot                     *bool   `form:"is_hot" json:"is_hot"`
	IsNew                     *bool   `form:"is_new" json:"is_new"`
	IsTable                   *bool   `form:"is_table" json:"is_table"`
	CategoryId                uint64  `form:"category_id" json:"category_id" `
	BrandId                   uint64  `form:"brand_id" json:"brand_id" `
	Name                      string  `form:"name" json:"name"`
	commonRequest.PageRequest         // 匿名嵌套直接展开
}
