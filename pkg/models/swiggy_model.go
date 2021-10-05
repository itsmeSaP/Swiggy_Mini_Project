package model

import (
	"time"
)

type User struct {
	ID       string `gorm:"default:uuid_generate_v4();primaryKey" json:"id"`
	Name     string `json:"title"`
	Number   string `json:"description"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type City struct {
	ID       string `gorm:"default:uuid_generate_v4();primaryKey" json:"id"`
	CityName string `json:"title"`
	State    string `json:"StateName"`
}
type Address struct {
	ID          string `gorm:"default:uuid_generate_v4();primaryKey" json:"id"`
	UserID      string `json:"user_id"`
	CityID      string `json:"city_id"`
	Zipcode     string `json:"zipcode"`
	AddressLine string `json:"address"`
	User        User   `gorm:"foreignKey:UserID"`
	City        City   `gorm:"foreignKey:CityID"`
}
type Restaurent struct {
	ID          string `gorm:"default:uuid_generate_v4();primaryKey" json:"id"`
	AddressLine string `json:"address"`
	CityID      string `json:"city_id"`
	Zipcode     string `json:"zipcode"`
	City        City   `gorm:"foreignKey:CityID"`
}
type FoodCategory struct {
	ID           string     `gorm:"default:uuid_generate_v4();primaryKey" json:"id"`
	RestaurentID string     `json:"restaurent_id"`
	CategoryName string     `json:"title"`
	Restaurent   Restaurent `gorm:"foreignKey:RestaurentID"`
}

type Menu struct {
	ID             string       `gorm:"default:uuid_generate_v4();primaryKey" json:"id"`
	RestaurentID   string       `json:"restaurent_id"`
	FoodCategoryID string       `json:"food_category"`
	Description    string       `json:"description"`
	Price          string       `json:"price"`
	Restaurent     Restaurent   `gorm:"foreignKey:RestaurentID"`
	FoodCategory   FoodCategory `gorm:"foreignKey:FoodCategoryID"`
}
type ItemsOrdered struct {
	ID       string `gorm:"default:uuid_generate_v4();primaryKey" json:"id"`
	MenuID   string `json:"menu_id"`
	Quantity string `json:"quantity"`
	Menu     Menu   `gorm:"foreignKey:MenuID"`
}
type Order struct {
	ID             string       `gorm:"default:uuid_generate_v4();primaryKey" json:"id"`
	UserID         string       `json:"user_id"`
	RestaurentID   string       `json:"restaurent_id"`
	ItemsOrderedID string       `json:"items_ordered_id"`
	AddressID      string       `json:"address_id"`
	OrderStatus    string       `json:"order_status"`
	OrderTime      time.Time    `json:"OrderTime"`
	DeliveryTime   time.Time    `json:"DeliveryTime"`
	TotalPrice     string       `json:"price"`
	User           User         `gorm:"foreignKey:UserID"`
	Restaurent     Restaurent   `gorm:"foreignKey:RestaurentID"`
	ItemsOrdered   ItemsOrdered `gorm:"foreignKey:ItemsOrderedID"`
	Address        Address      `gorm:"foreignKey:AddressID"`
}

type Payment struct {
	ID             string `gorm:"default:uuid_generate_v4();primaryKey" json:"id"`
	UserID         string `json:"user_id"`
	OrderID        string `json:"order_id"`
	AmountToBePaid string `json:"price"`
	PaymentStatus  string `json:"payment_status"`
	User           User   `gorm:"foreignKey:UserID"`
	Order          Order  `gorm:"foreignKey:OrderID"`
}
