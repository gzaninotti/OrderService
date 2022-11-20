package repository

import (
	"context"
)

type orderItemRepository interface {
	createNewOrderItem(ctx context.Context, orderItem entity.orderItem)
}
