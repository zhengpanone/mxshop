package forms

type AddressForm struct {
	Province     string `form:"province" json:"province" binding:"required"`
	City         string `form:"city" json:"city" binding:"required"`
	District     string `form:"district" json:"district" binding:"required"`
	Address      string `form:"address" json:"address" binding:"required"`
	SignerName   string `form:"signerName" json:"signerName" binding:"required"`
	SignerMobile string `form:"signerMobile" json:"signerMobile" binding:"required"`
}
