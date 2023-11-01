package interfaces

import (
	"context"

	"github.com/shion0625/backend/pkg/api/handler/request"
	"github.com/shion0625/backend/pkg/api/handler/response"
	"github.com/shion0625/backend/pkg/domain"
)

type OfferUseCase interface {

	// offer
	SaveOffer(ctx context.Context, offer request.Offer) error
	RemoveOffer(ctx context.Context, offerID uint) error
	FindAllOffers(ctx context.Context, pagination request.Pagination) ([]domain.Offer, error)

	// offer category
	SaveCategoryOffer(ctx context.Context, offerCategory request.OfferCategory) error
	FindAllCategoryOffers(ctx context.Context, pagination request.Pagination) ([]response.OfferCategory, error)
	RemoveCategoryOffer(ctx context.Context, categoryOfferID uint) error
	ChangeCategoryOffer(ctx context.Context, categoryOfferID, offerID uint) error

	// offer product
	SaveProductOffer(ctx context.Context, offerProduct domain.OfferProduct) error
	FindAllProductOffers(ctx context.Context, pagination request.Pagination) ([]response.OfferProduct, error)
	RemoveProductOffer(ctx context.Context, productOfferID uint) error
	ChangeProductOffer(ctx context.Context, productOfferID, offerID uint) error
}
