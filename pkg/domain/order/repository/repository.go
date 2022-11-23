package repository

import (
	"asdf/tcc/order-service/database"
	"asdf/tcc/order-service/pkg/domain/order/entity"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	collectionName = "order"
)

type OrderRepository interface {
	CreateNewOrder(ctx context.Context, order entity.Order) error
	EditOrderStatus(ctx context.Context, order entity.Order, filter primitive.M) error
	GetOrder(ctx context.Context, filter primitive.M) (entity.Order, error)
	GetUserOrders(ctx context.Context, filter primitive.M) ([]entity.Order, error)
	GetRestaurantOrders(ctx context.Context, filter primitive.M) ([]entity.Order, error)
}

type mongoOrderRepository struct {
	conn database.MongoConnection
}

func NewMongoOrderRepository(conn database.MongoConnection) *mongoOrderRepository {
	conn.Collection = conn.Database.Collection(collectionName)

	return &mongoOrderRepository{
		conn: conn,
	}
}

func (r *mongoOrderRepository) CreateNewOrder(ctx context.Context, order entity.Order) error {

	_, err := r.conn.Collection.InsertOne(ctx, order)
	if err != nil {
		return err
	}

	return nil
}

func (r *mongoOrderRepository) EditOrderStatus(ctx context.Context, order entity.Order, filter primitive.M) error {
	updateQuery := bson.M{
		"$set": bson.M{
			"status": order.Status,
		},
	}

	_, err := r.conn.Collection.UpdateOne(ctx, filter, updateQuery)
	if err != nil {
		return err
	}

	return nil
}

func (r *mongoOrderRepository) GetOrder(ctx context.Context, filter primitive.M) (entity.Order, error) {
	// verifica no serviço qual ID foi preenchido
	var foundOrder entity.Order

	if err := r.conn.Collection.FindOne(ctx, filter).Decode(foundOrder); err != nil {
		return entity.Order{}, nil
	}

	return foundOrder, nil

}

func (r *mongoOrderRepository) GetUserOrders(ctx context.Context, filter primitive.M) ([]entity.Order, error) {
	// verifica no serviço qual ID foi preenchido
	var userOrders []entity.Order

	filterCursor, err := r.conn.Collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	if err = filterCursor.All(ctx, &userOrders); err != nil {
		return nil, err
	}

	return userOrders, nil
}

func (r *mongoOrderRepository) GetRestaurantOrders(ctx context.Context, filter primitive.M) ([]entity.Order, error) {
	// verifica no serviço qual ID foi preenchido
	var orders []entity.Order

	filterCursor, err := r.conn.Collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	if err = filterCursor.All(ctx, &orders); err != nil {
		return nil, err
	}

	return orders, nil

}
