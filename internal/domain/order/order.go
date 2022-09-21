package order

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

type Order struct {
	ID          string
	OrderStatus orderStatus
	CustomerID  string
}
