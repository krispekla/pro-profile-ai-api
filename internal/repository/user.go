package repository

import (
	"database/sql"

	. "github.com/go-jet/jet/v2/postgres"
	"github.com/google/uuid"
	"github.com/krispekla/pro-profile-ai-api/.gen/postgres/auth/model"
	. "github.com/krispekla/pro-profile-ai-api/.gen/postgres/auth/table"
	_ "github.com/lib/pq"
)

type UserRepository interface {
	UpdateCustomerDetails() error
	Get(id uuid.UUID) (*model.Users, error)
}

type UserRepositoryImpl struct {
	Db *sql.DB
}

func NewUserRepositoryImpl(db *sql.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{
		Db: db,
	}
}

func (r *UserRepositoryImpl) Get(id uuid.UUID) (*model.Users, error) {
	stmt := SELECT(Users.AllColumns).FROM(Users).WHERE(Users.ID.EQ(UUID(id)))
	result := &model.Users{}
	err := stmt.Query(r.Db, result)
	return result, err
}

type UserCustomerInput struct {
	Id               uuid.UUID
	StripeCustomerID string
}

func (r *UserRepositoryImpl) UpdateCustomerDetails(usrInpt *UserCustomerInput) error {
	usr := model.Users{
		StripeCustomerID: &usrInpt.StripeCustomerID,
	}
	updateStmt := Users.UPDATE(Users.StripeCustomerID).MODEL(usr).WHERE(Users.ID.EQ(UUID(usrInpt.Id))).RETURNING(Users.AllColumns)
	dest := &model.Users{}
	err := updateStmt.Query(r.Db, dest)
	return err
}
