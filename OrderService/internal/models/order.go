package models

// OrderItem represents a single item in an order
type OrderItem struct {
	ProductID string  `bson:"productId" json:"productId"` // The ID of the product
	Quantity  int     `bson:"quantity" json:"quantity"`   // Quantity of the product in the order
	Price     float64 `bson:"price" json:"price"`         // Price per unit of the product
	Total     float64 `bson:"total" json:"total"`         // Total price for this item (Quantity * Price)
}

// Order represents the entire order which can have multiple items
type Order struct {
	BaseModel
	OrderID     string      `bson:"orderId" json:"orderId"`         // The ID of the order
	CustomerID  string      `bson:"customerId" json:"customerId"`   // The ID of the customer
	Status      string      `bson:"status" json:"status"`           // Order status (e.g., "pending", "completed")
	Items       []OrderItem `bson:"items" json:"items"`             // A list of OrderItems in this order
	TotalAmount float64     `bson:"totalAmount" json:"totalAmount"` // Total price of the order (sum of item totals)
}
