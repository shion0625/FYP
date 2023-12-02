package request

type PayOrder struct {
	UserID          string `json:"userId"    validate:"required"`
	AddressID       uint   `json:"addressId" validate:"required"`
	ProductItemInfo []ProductItemInfo
	TotalFee        uint `json:"totalFee" validate:"required,numeric"`
	VariationValue  *[]VariationValues
	PaymentMethodID uint `json:"paymentMethodId" validate:"required"`
}

type ProductItemInfo struct {
	ProductItemID uint `json:"productItemId" validate:"required"`
	Count         uint `json:"count"         validate:"required,numeric"`
}

type VariationValues struct {
	VariationID       uint `json:"variationId"       validate:"required,numeric"`
	VariationOptionID uint `json:"variationOptionId" validate:"required,gte=1"`
}
