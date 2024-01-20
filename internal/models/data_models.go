package models

type Customer struct {
	CustomerID string `json:"customer_id" bson:"customer_id"`
	FirstName  string `json:"first_name" bson:"first_name"`
	LastName   string `json:"last_name" bson:"last_name"`
	EmailID    string `json:"email" bson:"email"`
	Phone      string `json:"phone" bson:"phone"`
	Address    string `json:"address" bson:"address"`
}

type Product struct {
	ProductID   string  `json:"product_id" bson:"product_id"`
	ProductName string  `json:"product_name" bson:"name"`
	Price       float64 `json:"price" bson:"price"`
	VendorID    string  `json:"vendor_id" bson:"vendor_id"`
}

type Service struct {
	ServiceID   string  `json:"service_id" bson:"service_id"`
	ServiceName string  `json:"service_name" bson:"name"`
	Price       float64 `json:"price" bson:"price"`
}

type Vendor struct {
	VendorID string `json:"vendor_id" bson:"vendor_id"`
	Name     string `json:"name" bson:"name"`
	Contact  string `json:"contact" bson:"contact"`
	Phone    string `json:"phone" bson:"phone"`
	EmailID  string `json:"email" bson:"email"`
	Address  string `json:"address" bson:"address"`
}
