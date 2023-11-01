package interfaces

import (
	"github.com/shion0625/backend/pkg/api/handler/request"
	"github.com/shion0625/backend/pkg/domain"
)

type BrandRepository interface {
	IsExist(brand domain.Brand) (bool, error)
	Save(brand domain.Brand) (domain.Brand, error)
	Update(brand domain.Brand) error
	FindAll(pagination request.Pagination) ([]domain.Brand, error)
	FindOne(brandID uint) (domain.Brand, error)
	Delete(brandID uint) error
}
