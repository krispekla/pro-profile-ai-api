package repository

import (
	"database/sql"

	. "github.com/go-jet/jet/v2/postgres"
	"github.com/google/uuid"
	. "github.com/krispekla/pro-profile-ai-api/.gen/postgres/public/table"
	"github.com/krispekla/pro-profile-ai-api/types"
	_ "github.com/lib/pq"
)

type OrderRepository interface {
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
