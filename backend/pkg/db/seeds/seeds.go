package seeds

import (
	"fmt"

	"gorm.io/gorm"
)

const (
	UserCount    = 5
	ProductCount = 5
	AuthCount    = 5
)

// You can add the specified seed here.
func All(db *gorm.DB) []Seed {
	return []Seed{
		{
			Name: "RandomCreateUserDomain",
			Run: func(tx *gorm.DB) error {
				err := tx.Transaction(func(tx2 *gorm.DB) error {
					return CreateUsersDomain(tx2, UserCount)
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
	}
}
