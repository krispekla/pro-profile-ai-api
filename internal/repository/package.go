package repository

import (
	"database/sql"
	"errors"

	. "github.com/go-jet/jet/v2/postgres"
	"github.com/krispekla/pro-profile-ai-api/.gen/postgres/public/model"
	. "github.com/krispekla/pro-profile-ai-api/.gen/postgres/public/table"
	"github.com/krispekla/pro-profile-ai-api/types"
	_ "github.com/lib/pq"
)

type PackageRepository interface {
	Get() (*[]model.Package, error)
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
	).FROM(Package.FULL_JOIN(PackagePrice, PackagePrice.PackageID.EQ(Package.ID)))

	var result []types.PackageListingDTO

	err := stmt.Query(r.Db, &result)
	if err != nil {
		return nil, errors.New("error retrieving packages")
	}
	return &result, nil
}
