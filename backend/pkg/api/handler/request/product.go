package request

import "mime/multipart"

// for a new product.
type Product struct {
	Name            string `json:"name"        validate:"required,min=3,max=50"`
	Description     string `json:"description" validate:"required,min=10,max=300"`
	CategoryID      uint   `json:"categoryId"  validate:"required,number,gte=1"`
	BrandID         uint   `json:"brandId"     validate:"required,number,gte=1"`
	Price           uint   `json:"price"       validate:"required,numeric"`
	ImageFileHeader *multipart.FileHeader
}

type UpdateProduct struct {
	ID          uint   `json:"id"          validate:"required,number,gte=1"`
	Name        string `json:"name"        validate:"required,min=3,max=50"`
	Description string `json:"description" validate:"required,min=10,max=300"`
	CategoryID  uint   `json:"categoryId"  validate:"required,number,gte=1"`
	Price       uint   `json:"price"       validate:"required,numeric"`
	Image       string `json:"image"       validate:"required"`
}

// for a new productItem.
type ProductItem struct {
	Name               string                  `json:"name"               validate:"required,min=3,max=50"`
	Price              uint                    `json:"price"              validate:"required,numeric"`
	VariationOptionIDs []uint                  `json:"variationOptionIDs" validate:"required,gte=1"`
	QtyInStock         uint                    `json:"qtyInStock"         validate:"required,gte=1"`
	SKU                string                  `json:"-"`
	ImageFileHeaders   []*multipart.FileHeader `json:"imageFileHeaders"   validate:"required,min=0"`
}

type Variation struct {
	Names []string `json:"names" validate:"required,dive,min=1"`
}

type VariationOption struct {
	Values []string `json:"values" validate:"required,dive,min=1"`
}

type Category struct {
	Name string `json:"name" validate:"required,min=1,max=50"`
}

type Brand struct {
	Name string `json:"name" validate:"required,min=3,max=25"`
}
