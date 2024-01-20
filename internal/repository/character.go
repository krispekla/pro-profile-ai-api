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

type UpdateCharacterInput struct {
	ID        int64           `json:"id"`
	Name      string          `json:"name"`
	Gender    model.Gender    `json:"gender"`
	HairColor model.HairColor `json:"hairColor"`
	EyeColor  model.EyeColor  `json:"eyeColor"`
	Ethnicity model.Ethnicity `json:"ethnicity"`
	Age       model.Age       `json:"age"`
}

type CreateCharacterDTO struct {
	ID        int             `alias:"character.id" json:"id"`
	Name      string          `alias:"character.name" json:"name"`
	Gender    model.Gender    `alias:"character.gender" json:"gender"`
	HairColor model.HairColor `alias:"character.hair_color" json:"hairColor"`
	EyeColor  model.EyeColor  `alias:"character.eye_color" json:"eyeColor"`
	Ethnicity model.Ethnicity `alias:"character.ethnicity" json:"ethnicity"`
	Age       model.Age       `alias:"character.age" json:"age"`
}

type CharacterRepository interface {
	GetCharacter(id uuid.UUID) (*[]CharacterDTO, error)
	CreateCharacter(usrId uuid.UUID) (*CreateCharacterDTO, error)
	UpdateCharacter(inp *UpdateCharacterInput) (*CreateCharacterDTO, error)
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

func (r *CharacterRepositoryImpl) CreateCharacter(usrId uuid.UUID) (*CreateCharacterDTO, error) {
	stmt := Character.INSERT(
		Character.UserID,
	).VALUES(UUID(usrId)).RETURNING(Character.AllColumns.Except(Character.UserID))
	res := &CreateCharacterDTO{}
	err := stmt.Query(r.Db, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *CharacterRepositoryImpl) UpdateCharacter(inp *UpdateCharacterInput) (*CreateCharacterDTO, error) {
	// Setting default values if no value is present
	if inp.Gender == "" {
		inp.Gender = model.Gender_Other
	}
	if inp.HairColor == "" {
		inp.HairColor = model.HairColor_Other
	}
	if inp.EyeColor == "" {
		inp.EyeColor = model.EyeColor_Brown
	}
	if inp.Ethnicity == "" {
		inp.Ethnicity = model.Ethnicity_Other
	}
	if inp.Age == "" {
		inp.Age = model.Age_Adult
	}

	chr := model.Character{
		Name:      inp.Name,
		Gender:    inp.Gender,
		HairColor: inp.HairColor,
		EyeColor:  inp.EyeColor,
		Ethnicity: inp.Ethnicity,
		Age:       inp.Age,
	}
	stmt := Character.UPDATE(
		Character.AllColumns.Except(Character.ID, Character.Created),
	).MODEL(chr).WHERE(Character.ID.EQ(Int(inp.ID))).RETURNING(Character.AllColumns.Except(Character.UserID))
	resChr := &CreateCharacterDTO{}
	err := stmt.Query(r.Db, resChr)
	if err != nil {
		return nil, err
	}
	return resChr, nil
}
