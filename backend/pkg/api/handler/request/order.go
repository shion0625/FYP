package request

type PayOrder struct {
	UserID          string `binding:"required" json:"userId"`
	AddressID       uint   `binding:"required" json:"addressId"`
	ProductItemInfo []ProductItemInfo
	TotalFee        uint `binding:"required,numeric" json:"totalFee"`
	VariationValue  *[]VariationValues
	PaymentMethodID uint `binding:"required" json:"paymentMethodId"`
}

type ProductItemInfo struct {
	ProductItemID uint `binding:"required"         json:"productItemId"`
	Count         uint `binding:"required,numeric" json:"count"`
}

type VariationValues struct {
	VariationID       uint `binding:"required,numeric" json:"variationId"`
	VariationOptionID uint `binding:"required,gte=1"   json:"variationOptionId"`
}
