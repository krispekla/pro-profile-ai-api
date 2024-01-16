package repository

import (
	"database/sql"

	. "github.com/go-jet/jet/v2/postgres"
	"github.com/google/uuid"
	"github.com/krispekla/pro-profile-ai-api/.gen/postgres/public/model"
	. "github.com/krispekla/pro-profile-ai-api/.gen/postgres/public/table"
	"github.com/krispekla/pro-profile-ai-api/types"
	_ "github.com/lib/pq"
)

type OrderRepository interface {
	CreateOrder(inp *CreateOrderInput) (*model.PackageOrder, error)
	GetAllOrders(id uuid.UUID) (*[]types.OrderAllDTO, error)
}

type OrderRepositoryImpl struct {
	db *sql.DB
}

func NewOrderRepositoryImpl(db *sql.DB) *OrderRepositoryImpl {
	return &OrderRepositoryImpl{
		db: db,
	}
}

func (r *OrderRepositoryImpl) GetAllOrders(id uuid.UUID) (*[]types.OrderAllDTO, error) {
	stmt := SELECT(
		PackageOrder.AllColumns,
	).FROM(PackageOrder).WHERE(PackageOrder.UserID.EQ(UUID(id)))

	var result []types.OrderAllDTO

	err := stmt.Query(r.db, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

type CreateOrderInput struct {
	PaymentIntentId *string
	Amount          *int64
	Currency        string
	UserId          *uuid.UUID
	PackagePrices   *[]model.PackagePrice
	CouponId        *string
}

func (r *OrderRepositoryImpl) CreateOrder(inp *CreateOrderInput) (*model.PackageOrder, error) {
	newOrderStmt := PackageOrder.INSERT(
		PackageOrder.StripePaymentIntentID,
		PackageOrder.TotalAmount,
		PackageOrder.Currency,
		PackageOrder.UserID,
		PackageOrder.CouponID,
	).VALUES(
		inp.PaymentIntentId,
		inp.Amount,
		inp.Currency,
		&inp.UserId,
	).RETURNING(
		PackageOrder.AllColumns,
	)
	var newOrder model.PackageOrder

	err := newOrderStmt.Query(r.db, &newOrder)
	if err != nil {
		return nil, err
	}

	orderItems := make([]model.PackageOrderItem, len(*inp.PackagePrices))
	totalPrice := int32(0)
	for i, pprice := range *inp.PackagePrices {
		orderItems[i] = model.PackageOrderItem{
			PackageOrderID: newOrder.ID,
			PackageID:      pprice.PackageID,
		}
		totalPrice += pprice.Amount
	}
	newOrder.TotalAmount = totalPrice

	newOrderItemsStmt := PackageOrderItem.INSERT(
		PackageOrderItem.PackageOrderID,
		PackageOrderItem.PackageID,
	).MODEL(&orderItems).RETURNING(PackageOrderItem.AllColumns)

	err = newOrderItemsStmt.Query(r.db, &orderItems)
	if err != nil {
		return nil, err
	}

	uNewOrder := &model.PackageOrder{}
	updateNewOrderStmt :=
		PackageOrder.UPDATE(
			PackageOrder.TotalAmount,
		).SET(&totalPrice).WHERE(
			PackageOrder.ID.EQ(Int32(newOrder.ID)),
		).RETURNING(PackageOrder.AllColumns)
	err = updateNewOrderStmt.Query(r.db, uNewOrder)
	if err != nil {
		return nil, err
	}
	return uNewOrder, nil
}
