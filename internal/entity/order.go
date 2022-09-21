package entity

const (
	Accepted       string = "ACCEPTED"
	Aproved        string = "APPROVED"
	Cancelled      string = "CANCELLED"
	Preparing      string = "PREPARING"
	ReadyForPickup string = "READY_FOR_PICKUP"
	PickedUp       string = "PICKED_UP"
	Delivered      string = "DELIVERED"
)

type Order struct {
	OrderState string
	Courier    string
}
