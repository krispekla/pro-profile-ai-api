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
	UpdateOrder(chkSessId string, paymInteId string, status string) (*model.PackageOrder, error)
	CreateGeneratedPackage(orderId int) (*[]model.GeneratedPackage, error)
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
	Amount        *int64
	Currency      string
	UserId        *uuid.UUID
	PackagePrices *[]model.PackagePrice
	CouponId      *string
	CheckoutId    *string
}

func (r *OrderRepositoryImpl) CreateOrder(inp *CreateOrderInput) (*model.PackageOrder, error) {
	newOrderStmt := PackageOrder.INSERT(
		PackageOrder.StripeCheckoutSessionID,
		PackageOrder.TotalAmount,
		PackageOrder.Currency,
		PackageOrder.UserID,
	).VALUES(
		inp.CheckoutId,
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

	var orderItems = &[]model.PackageOrderItem{}
	totalPrice := int32(0)
	for _, pprice := range *inp.PackagePrices {
		*orderItems = append(*orderItems, model.PackageOrderItem{
			PackageOrderID: newOrder.ID,
			PackageID:      pprice.PackageID,
		})
		totalPrice += pprice.Amount
	}
	newOrder.TotalAmount = totalPrice

	newOrderItemsStmt := PackageOrderItem.INSERT(
		PackageOrderItem.PackageOrderID,
		PackageOrderItem.PackageID,
	).MODELS(orderItems).RETURNING(PackageOrderItem.AllColumns)

	_, err = newOrderItemsStmt.Exec(r.db)
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

func (r *OrderRepositoryImpl) UpdateOrder(chkSessId string, paymInteId string, status string) (*model.PackageOrder, error) {
	return nil, nil
}

func (r *OrderRepositoryImpl) CreateGeneratedPackage(orderId int) (*[]model.GeneratedPackage, error) {
	return nil, nil
}
