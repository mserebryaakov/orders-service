package order

import (
	"context"
	"orders-service/internal/domain/order"
	repository "orders-service/internal/domain/order"
)

// Сервис взаимодействия с заказами
type OrderUseCase struct {
	repo repository.OrderRepository
}

// Создание сервиса взаимодействия с заказами
func NewOrderService(or repository.OrderRepository) *OrderUseCase {
	return &OrderUseCase{
		repo: or,
	}
}

// Usecase создания заказа
func (s *OrderUseCase) CreateItem(ctx context.Context, order order.Order) (string, error) {
	id, err := s.repo.Create(ctx, order)
	if err != nil {
		return "", err
	}

	return id, nil
}

// Usecase поиска заказа
func (s *OrderUseCase) FindOne(ctx context.Context, id string) (o order.Order, err error) {
	o, err = s.repo.FindOne(ctx, id)
	if err != nil {
		return o, err
	}

	return
}

// Usecase обновления заказа
func (s *OrderUseCase) Update(ctx context.Context, order order.Order) (err error) {
	err = s.repo.Update(ctx, order)
	if err != nil {
		return err
	}

	return
}

// Usecase удаления заказа
func (s *OrderUseCase) Delete(ctx context.Context, id string) (err error) {
	err = s.repo.Delete(ctx, id)
	if err != nil {
		return err
	}

	return
}
