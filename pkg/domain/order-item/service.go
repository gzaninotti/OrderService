package orderitem

import (
	"asdf/tcc/order-service/pkg/domain/order-item/entity"
	"asdf/tcc/order-service/pkg/domain/order-item/repository"
	uuidService "asdf/tcc/order-service/pkg/lib/uuid"
	"context"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderItemService interface {
	CreateNewOrderItem(ctx context.Context, orderItem entity.OrderItem) (entity.OrderItem, error)
	EditOrderItemStatus(ctx context.Context, orderItem entity.OrderItem) (entity.OrderItem, error)
	GetOrderItem(ctx context.Context, orderItem entity.OrderItem) (entity.OrderItem, error)
	ListOrderItens(ctx context.Context, orderItem entity.OrderItem) ([]entity.OrderItem, error)
}

type orderItemService struct {
	uuidService         uuidService.Service
	orderItemRepository repository.OrderItemRepository
}

func NewOrderItemService(
	uuidService uuidService.Service,
	orderItemrepository repository.OrderItemRepository,
) OrderItemService {
	return &orderItemService{
		uuidService:         uuidService,
		orderItemRepository: orderItemrepository,
	}
}

func (s *orderItemService) CreateNewOrderItem(ctx context.Context, orderItem entity.OrderItem) (entity.OrderItem, error) {
	// If needs a validation here it is supposed to be
	// err := s.orderItemValidator.ValidadeOnCreate(orderItem)
	// if err != nil { retorna }

	orderItem.ID = s.uuidService.Generate()
	orderItem.Status = 0
	orderItem.Price = calculateOrderItemPrice(orderItem)

	err := s.orderItemRepository.CreateNewOrderItem(ctx, orderItem)
	if err != nil {
		return entity.OrderItem{}, err
	}

	return orderItem, nil

}

func (s *orderItemService) EditOrderItemStatus(ctx context.Context, orderItem entity.OrderItem) (entity.OrderItem, error) {
	// If needs a validation here it is supposed to be
	// err := s.orderItemValidator.ValidadeOnUpdate(orderItem)
	// if err != nil { retorna }
	filter := bson.M{}
	filter = makeQueryFilter(orderItem)

	if err := s.orderItemRepository.EditOrderItemStatus(ctx, orderItem, filter); err != nil {
		return entity.OrderItem{}, err
	}

	return orderItem, nil

}

func (s *orderItemService) GetOrderItem(ctx context.Context, orderItem entity.OrderItem) (entity.OrderItem, error) {
	filter := bson.M{}
	filter = makeQueryFilter(orderItem)

	orderItem, err := s.orderItemRepository.GetOrderItem(ctx, filter)
	if err != nil {
		return entity.OrderItem{}, nil
	}

	return orderItem, nil
}

func (s *orderItemService) ListOrderItens(ctx context.Context, orderItem entity.OrderItem) ([]entity.OrderItem, error) {
	filter := bson.M{}
	filter = makeQueryFilter(orderItem)

	orderItens, err := s.orderItemRepository.ListOrderItens(ctx, filter)
	if err != nil {
		return nil, err
	}

	return orderItens, nil
}

func calculateOrderItemPrice(orderItem entity.OrderItem) float64 {
	var price float64
	quantity := orderItem.Quantity
	unitPrice := orderItem.UnitPrice

	price = unitPrice * float64(quantity)

	return price
}

func makeQueryFilter(order entity.OrderItem) primitive.M {
	filter := bson.M{}

	if order.ID != uuid.Nil {
		filter["id"] = bson.M{"_id": order.ID}
	} else if order.OrderID != uuid.Nil {
		filter["user_id"] = bson.M{"user_id": order.OrderID}
	}

	return filter
}
