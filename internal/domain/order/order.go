package order

// Статусы заказа
type orderStatus string

const (
	Accepted       orderStatus = "ACCEPTED"
	Aproved        orderStatus = "APPROVED"
	Cancelled      orderStatus = "CANCELLED"
	Preparing      orderStatus = "PREPARING"
	ReadyForPickup orderStatus = "READY_FOR_PICKUP"
	PickedUp       orderStatus = "PICKED_UP"
	Delivered      orderStatus = "DELIVERED"
)

// Модель заказа
type Order struct {
	ID          string      `json:"id" bson:"_id,omitempty"`
	OrderStatus orderStatus `json:"order_status" bson:"order_status" binding:"required"`
	CustomerID  string      `json:"customer_id" bson:"customer_id" binding:"required"`
}
