package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID              primitive.ObjectID `json:"_id" bson:"id"`
	First_name      *string            `json:"first_name" validate:"required, min=2, max=30"`
	Last_name       *string            `json:"last_name" validate:"requirde, min=2, max=30"`
	Password        *string            `json:"password" validate:"required, min=6"`
	Email           *string            `json:"email" validate:"email, required"`
	Phone           *string            `json:"phone" validate:"required"`
	Token           *string            `json:"token"`
	Refresh_Token   *string            `json:"resresh_token"`
	Created_At      time.Time          `json:"creatd_at"`
	Updated_At      time.Time          `json:"upadted_at"`
	User_ID         string             `json:"user_id"`
	UserCart        []ProductUser      `json:"user_cart" bson:"user_cart"`
	Address_Details []Address          `json:"address_details" bson:"address_details"`
	Order_Status    []Order            `json:"order_status" bson:"order_status"`
}

type Product struct {
	Product_ID   primitive.ObjectID `json:"product_id" bson:"product_id"`
	Product_Name *string            `json:"product_name"`
	Price        uint64             `json:"price" bson:"price"`
	Rating       *uint64            `json:"rating" bson:"rating"`
	Image        *string            `json:"image" bson:"image"`
}
type ProductUser struct {
	Product_ID   primitive.ObjectID `josn:"product_id" bson:"product_id"`
	Product_Name *string            `json:"product_name"`
	Price        uint64             `json:"price"`
	Rating       *uint              `json:"rating"`
	Image        *string            `json:"image"`
}
type Order struct {
	Order_ID       primitive.ObjectID `json:"order_id" bson:"order_id"`
	Order_Card     []ProductUser      `json:"order_card" bson:"order_card"`
	Ordered_At     time.Time          `json:"ordered_at" bson:"ordered_at"`
	Price          int                `json:"price" bson:"price"`
	Discount       *int               `json:"discount" bson:"discount"`
	Payment_Method Payment            `json:"payment_method" bson:"payment_method"`
}
type Payment struct {
	Digital bool `json:"digital"`
	COD     bool `json:"cod"`
}
type Address struct {
	Address_id primitive.ObjectID `json:"address_id" bson:"address_id"`
	House      *string            `json:"house_name" bson:"house_name"`
	Street     *string            `json:"street_name" bson:"street_name"`
	City       *string            `json:"city_name" bson:"city_name"`
	Pincode    *string            `json:"pin_code" bson:"pin_code"`
}
