package repository

import (
	"asdf/tcc/order-service/database"
	"asdf/tcc/order-service/pkg/domain/order-item/entity"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	collectionName = "orderItem"
)

type OrderItemRepository interface {
	CreateNewOrderItem(ctx context.Context, orderItem entity.OrderItem) error
	EditOrderItemStatus(ctx context.Context, orderItem entity.OrderItem, filter primitive.M) error
	GetOrderItem(ctx context.Context, filter primitive.M) (entity.OrderItem, error)
	ListOrderItens(ctx context.Context, filter primitive.M) ([]entity.OrderItem, error)
}

type mongoOrderItemRepository struct {
	conn database.MongoConnection
}

func NewMongoOrderItemRepository(conn database.MongoConnection) *mongoOrderItemRepository {
	conn.Collection = conn.Database.Collection(collectionName)

	return &mongoOrderItemRepository{
		conn: conn,
	}
}

func (r *mongoOrderItemRepository) CreateNewOrderItem(ctx context.Context, orderItem entity.OrderItem) error {

	_, err := r.conn.Collection.InsertOne(ctx, orderItem)
	if err != nil {
		return err
	}

	return nil
}

func (r *mongoOrderItemRepository) EditOrderItemStatus(ctx context.Context, orderItem entity.OrderItem, filter primitive.M) error {
	updateQuery := bson.M{
		"$set": bson.M{
			"status": orderItem.Status,
		},
	}

	_, err := r.conn.Collection.UpdateOne(ctx, filter, updateQuery)
	if err != nil {
		return err
	}

	return nil
}

func (r *mongoOrderItemRepository) GetOrderItem(ctx context.Context, filter primitive.M) (entity.OrderItem, error) {
	var orderItem entity.OrderItem

	if err := r.conn.Collection.FindOne(ctx, filter).Decode(orderItem); err != nil {
		return entity.OrderItem{}, nil
	}

	return orderItem, nil

}

func (r *mongoOrderItemRepository) ListOrderItens(ctx context.Context, filter primitive.M) ([]entity.OrderItem, error) {
	var orderItens []entity.OrderItem

	filterCursor, err := r.conn.Collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	if err = filterCursor.All(ctx, &orderItens); err != nil {
		return nil, err
	}

	return orderItens, nil

}
