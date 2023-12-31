package seeds

import (
	"fmt"

	"github.com/shion0625/FYP/backend/pkg/config"
	"gorm.io/gorm"
)

const (
	UserCount    = 10
	ProductCount = 10
	AuthCount    = 10
	OrderCount   = 10
)

// You can add the specified seed here.
func All(db *gorm.DB, cfg *config.Config) []Seed {
	return []Seed{
		{
			Name: "RandomCreateUserDomain",
			Run: func(tx *gorm.DB) error {
				err := tx.Transaction(func(tx2 *gorm.DB) error {
					return CreateUsersDomain(tx2, UserCount, cfg)
				})
				if err != nil {
					return fmt.Errorf("failed to create user domain: %w", err)
				}

				return nil
			},
		},
		{
			Name: "RandomCreateProductDomain",
			Run: func(tx *gorm.DB) error {
				err := tx.Transaction(func(tx2 *gorm.DB) error {
					return CreateProductsDomain(tx2, ProductCount)
				})
				if err != nil {
					return fmt.Errorf("failed to create product domain: %w", err)
				}

				return nil
			},
		},
		{
			Name: "RandomCreateOrderDomain",
			Run: func(tx *gorm.DB) error {
				err := tx.Transaction(func(tx2 *gorm.DB) error {
					return CreateOrdersDomain(tx2, OrderCount)
				})
				if err != nil {
					return fmt.Errorf("failed to create product domain: %w", err)
				}

				return nil
			},
		},
	}
}
