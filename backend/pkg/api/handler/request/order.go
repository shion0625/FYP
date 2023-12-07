package request

type PayOrder struct {
	UserID          string            `json:"userId"          validate:"required"`
	AddressID       uint              `json:"addressId"       validate:"required,number,gte=1"`
	ProductItemInfo []ProductItemInfo `json:"productItemInfo" validate:"required,min=1"`
	TotalFee        uint              `json:"totalFee"        validate:"required,numeric"`
	PaymentMethodID uint              `json:"paymentMethodId" validate:"required,number,gte=1"`
}

type ProductItemInfo struct {
	ProductItemID   uint               `json:"productItemId"   validate:"required,number,gte=1"`
	Count           uint               `json:"count"           validate:"required,number"`
	VariationValues *[]VariationValues `json:"variationValues"`
}

type VariationValues struct {
	VariationID       uint   `json:"variationId"       validate:"required,number,gte=1"`
	Name              string `json:"name"`
	VariationOptionID uint   `json:"variationOptionId" validate:"required,number,gte=1"`
	Value             string `json:"value"`
}
