package repository

import (
	"database/sql"
	"errors"

	. "github.com/go-jet/jet/v2/postgres"
	"github.com/google/uuid"
	authTable "github.com/krispekla/pro-profile-ai-api/.gen/postgres/auth/table"
	"github.com/krispekla/pro-profile-ai-api/.gen/postgres/public/model"
	. "github.com/krispekla/pro-profile-ai-api/.gen/postgres/public/table"
	_ "github.com/lib/pq"
)

// TODO: Move to types
type CharacterDTO struct {
	model.Character
	Email string `db:"email"`
}

type CharacterRepository interface {
	GetCharacter(id uuid.UUID) (*[]CharacterDTO, error)
}

type CharacterRepositoryImpl struct {
	Db *sql.DB
}

func NewCharacterRepositoryImpl(db *sql.DB) *CharacterRepositoryImpl {
	return &CharacterRepositoryImpl{
		Db: db,
	}
}

func (r *CharacterRepositoryImpl) Get(id uuid.UUID) (*[]CharacterDTO, error) {
	stmt := SELECT(
		Character.AllColumns,
		authTable.Users.Email.AS("Email"),
	).FROM(Character.LEFT_JOIN(
		authTable.Users, Character.UserID.EQ(UUID(id))))

	var result []CharacterDTO
	err := stmt.Query(r.Db, &result)
	if err != nil {
		return nil, errors.New("error retrieving characters")
	}
	return &result, nil
}
