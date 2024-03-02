package usecase

import (
	"github.com/ebdonato/go-clean-arch/internal/entity"
	"github.com/ebdonato/go-clean-arch/pkg/events"
)

// DTO - entradas e saida de dados
type OrdersInputDTO struct{}

type OrdersOutputDTO struct {
	Orders []entity.Order
}

// Os componentes do nosso use case.
type GetOrdersUseCase struct {
	OrderRepository entity.OrderRepositoryInterface // Interface para falar com o BD.
	OrderCreated    events.EventInterface           // Interface de ordem criada - evento.
	EventDispatcher events.EventDispatcherInterface // Interface para disparo do evento.
}

// Criação dos componentes / usecase
func NewGetOrdersUseCase(
	OrderRepository entity.OrderRepositoryInterface,
	OrderCreated events.EventInterface,
	EventDispatcher events.EventDispatcherInterface,
) *GetOrdersUseCase {
	return &GetOrdersUseCase{
		OrderRepository: OrderRepository,
		OrderCreated:    OrderCreated,
		EventDispatcher: EventDispatcher,
	}
}

func (c *GetOrdersUseCase) Execute() (OrdersOutputDTO, error) {
	orders, err := c.OrderRepository.GetAll()
	if err != nil {
		return OrdersOutputDTO{}, err
	}

	dto := OrdersOutputDTO{
		Orders: orders,
	}

	c.OrderCreated.SetPayload(dto)             // interface
	c.EventDispatcher.Dispatch(c.OrderCreated) // interface

	return dto, nil
}
