package order

import (
	"asdf/tcc/order-service/pkg/domain/order/entity"
	"asdf/tcc/order-service/pkg/domain/order/repository"
	uuidService "asdf/tcc/order-service/pkg/lib/uuid"
	"context"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderService interface {
	CreateNewOrder(ctx context.Context, order entity.Order) (entity.Order, error)
	EditOrderStatus(ctx context.Context, order entity.Order) (entity.Order, error)
	GetOrder(ctx context.Context, order entity.Order) (entity.Order, error)
	GetUserOrders(ctx context.Context, order entity.Order) ([]entity.Order, error)
	GetRestaurantOrders(ctx context.Context, order entity.Order) ([]entity.Order, error)
}

type orderService struct {
	uuidService     uuidService.Service
	orderRepository repository.OrderRepository
}

func NewOrderService(
	uuidService uuidService.Service,
	orderrepository repository.OrderRepository,
) OrderService {
	return &orderService{
		uuidService:     uuidService,
		orderRepository: orderrepository,
	}
}

func (s *orderService) CreateNewOrder(ctx context.Context, order entity.Order) (entity.Order, error) {
	// If needs a validation here it is supposed to be
	// err := s.orderValidator.ValidadeOnCreate(order)
	// if err != nil { retorna }

	// preenchido com o UserID
	order.ID = s.uuidService.Generate()
	order.Status = 0

	err := s.orderRepository.CreateNewOrder(ctx, order)
	if err != nil {
		return entity.Order{}, err
	}

	return order, nil

}

func (s *orderService) EditOrderStatus(ctx context.Context, order entity.Order) (entity.Order, error) {
	// If needs a validation here it is supposed to be
	// err := s.orderValidator.ValidadeOnUpdate(order)
	// if err != nil { retorna }

	// Make filter based on which ID field is filled
	filter := bson.M{}
	filter = makeQueryFilter(order)

	if err := s.orderRepository.EditOrderStatus(ctx, order, filter); err != nil {
		return entity.Order{}, err
	}

	return order, nil

}

func (s *orderService) GetOrder(ctx context.Context, order entity.Order) (entity.Order, error) {
	filter := bson.M{}
	filter = makeQueryFilter(order)

	order, err := s.orderRepository.GetOrder(ctx, filter)
	if err != nil {
		return entity.Order{}, nil
	}

	return order, nil
}

func (s *orderService) GetUserOrders(ctx context.Context, order entity.Order) ([]entity.Order, error) {
	filter := bson.M{}
	filter = makeQueryFilter(order)

	orderItens, err := s.orderRepository.GetUserOrders(ctx, filter)
	if err != nil {
		return nil, err
	}

	return orderItens, nil
}

func (s *orderService) GetRestaurantOrders(ctx context.Context, order entity.Order) ([]entity.Order, error) {
	filter := bson.M{}
	filter = makeQueryFilter(order)

	orderItens, err := s.orderRepository.GetRestaurantOrders(ctx, filter)
	if err != nil {
		return nil, err
	}

	return orderItens, nil
}

func makeQueryFilter(order entity.Order) primitive.M {
	filter := bson.M{}

	if order.ID != uuid.Nil {
		filter["id"] = bson.M{"_id": order.ID}
	} else if order.UserID != uuid.Nil {
		filter["user_id"] = bson.M{"user_id": order.UserID}
	} else if order.UserID != uuid.Nil {
		filter["table_id"] = bson.M{"table_id": order.TableID}
	}

	return filter
}
