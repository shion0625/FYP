package response

type Order struct {
	UserID          string            `json:"userId"`
	ShopOrderId     string            `json:"shopOrderId"`
	ProductItemInfo []ProductItemInfo `json:"productItemInfo"`
	Address         Address           `json:"address"`
	TotalFee        uint              `json:"totalFee"`
	PaymentMethod   PaymentMethod     `json:"paymentMethod"`
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
