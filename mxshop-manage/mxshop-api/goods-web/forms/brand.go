package forms

// BrandForm 品牌分类
type BrandForm struct {
	Name string `form:"name" json:"name" binding:"required,min=3,max=10"` // 品牌名称
	Logo string `form:"logo" json:"logo" binding:"url"`                   // 品牌logo
}

type CategoryBrandForm struct {
	CategoryId int `form:"category_id" json:"category_id" binding:"required"` // 商品分类ID
	BrandId    int `form:"brand_id" json:"brand_id" binding:"required"`       // 品牌ID
}
