package request

import "mime/multipart"

// for a new product.
type Product struct {
	Name            string `binding:"required,min=3,max=50"   json:"name"`
	Description     string `binding:"required,min=10,max=100" json:"description"`
	CategoryID      uint   `binding:"required"                json:"categoryId"`
	BrandID         uint   `binding:"required"                json:"brandId"`
	Price           uint   `binding:"required,numeric"        json:"price"`
	ImageFileHeader *multipart.FileHeader
}

type UpdateProduct struct {
	ID          uint   `binding:"required"                json:"id"`
	Name        string `binding:"required,min=3,max=50"   json:"name"`
	Description string `binding:"required,min=10,max=100" json:"description"`
	CategoryID  uint   `binding:"required"                json:"categoryId"`
	Price       uint   `binding:"required,numeric"        json:"price"`
	Image       string `binding:"required"                json:"image"`
}

// for a new productItem.
type ProductItem struct {
	Name               string                  `binding:"required,min=3,max=50"   json:"name"`
	Price              uint                    `binding:"required,min=1" json:"price"`
	VariationOptionIDs []uint                  `binding:"required,gte=1" json:"variationOptionIDs"`
	QtyInStock         uint                    `binding:"required,min=1" json:"qtyInStock"`
	SKU                string                  `json:"-"`
	ImageFileHeaders   []*multipart.FileHeader `binding:"required,gte=1" json:"imageFileHeaders"`
}

type Variation struct {
	Names []string `binding:"required,dive,min=1" json:"names"`
}

type VariationOption struct {
	Values []string `binding:"required,dive,min=1" json:"values"`
}

type Category struct {
	Name string `binding:"required" json:"name"`
}

type Brand struct {
	Name string `binding:"required,min=3,max=25" json:"name"`
}
