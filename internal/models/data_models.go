package models

type Customer struct {
	CustomerID string `json:"customer_id" bson:"customer_id"`
	FirstName  string `json:"first_name" bson:"first_name"`
	LastName   string `json:"last_name" bson:"last_name"`
	EmailID    string `json:"email" bson:"email"`
	Phone      string `json:"phone" bson:"phone"`
	Address    string `json:"address" bson:"address"`
}
