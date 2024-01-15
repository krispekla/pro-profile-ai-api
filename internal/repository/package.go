package repository

import (
	"database/sql"
	"errors"

	. "github.com/go-jet/jet/v2/postgres"
	"github.com/google/uuid"
	"github.com/krispekla/pro-profile-ai-api/.gen/postgres/public/model"
	. "github.com/krispekla/pro-profile-ai-api/.gen/postgres/public/table"
	"github.com/krispekla/pro-profile-ai-api/types"
	_ "github.com/lib/pq"
)

type PackageRepository interface {
	GetListing() (*[]types.PackageListingDTO, error)
	GetGeneratedPackages(usrId uuid.UUID) (*[]types.PackageGeneratedDTO, error)
	GetPackagePrice(productId *[]string) (*[]model.PackagePrice, error)
}

type PackageRepositoryImpl struct {
	Db *sql.DB
}

func NewPackageRepositoryImpl(db *sql.DB) *PackageRepositoryImpl {
	return &PackageRepositoryImpl{
		Db: db,
	}
}

func (r *PackageRepositoryImpl) GetListing() (*[]types.PackageListingDTO, error) {
	stmt := SELECT(
		Package.AllColumns,
		PackagePrice.AllColumns,
		PackageExampleImg.AllColumns.Except(PackageExampleImg.PackageID, PackageExampleImg.ID),
	).FROM(Package.LEFT_JOIN(PackagePrice, PackagePrice.PackageID.EQ(Package.ID)).LEFT_JOIN(PackageExampleImg, Package.ID.EQ(PackageExampleImg.PackageID)))

	var result []types.PackageListingDTO

	err := stmt.Query(r.Db, &result)
	if err != nil {
		return nil, errors.New("error retrieving packages")
	}
	return &result, nil
}

func (r *PackageRepositoryImpl) GetGeneratedPackages(usrId uuid.UUID) (*[]types.PackageGeneratedDTO, error) {
	stmt := SELECT(
		GeneratedPackage.AllColumns,
		PackageOrderItem.PackageID,
	).FROM(
		GeneratedPackage.LEFT_JOIN(PackageOrderItem, PackageOrderItem.ID.EQ(GeneratedPackage.PackageOrderItemID)).LEFT_JOIN(PackageOrder, PackageOrder.ID.EQ(PackageOrderItem.PackageOrderID)),
	).WHERE(
		PackageOrder.UserID.EQ(UUID(usrId)),
	)

	var result []types.PackageGeneratedDTO

	err := stmt.Query(r.Db, &result)
	if err != nil {
		return nil, errors.New("error retrieving packages")
	}
	return &result, nil
}

func (r *PackageRepositoryImpl) GetPackagePrice(productIds *[]string) (*[]model.PackagePrice, error) {
	var pIds []Expression
	for _, pId := range *productIds {
		pIds = append(pIds, String(pId))
	}
	stmt := SELECT(
		PackagePrice.AllColumns,
	).FROM(PackagePrice).WHERE(PackagePrice.StripeProductID.IN(pIds...))

	var result []model.PackagePrice

	err := stmt.Query(r.Db, &result)
	if err != nil {
		return nil, errors.New("error retrieving packages")
	}
	return &result, nil
}
