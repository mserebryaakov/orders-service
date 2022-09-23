package order

import "context"

type OrderRepository interface {
	Create(ctx context.Context, order Order) (string, error)
	FindOne(ctx context.Context, id string) (Order, error)
	Update(ctx context.Context, order Order) error
	Delete(ctx context.Context, id string) error
}
