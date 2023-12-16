//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package table

import (
	"github.com/go-jet/jet/v2/postgres"
)

var MfaFactors = newMfaFactorsTable("auth", "mfa_factors", "")

type mfaFactorsTable struct {
	postgres.Table

	// Columns
	ID           postgres.ColumnString
	UserID       postgres.ColumnString
	FriendlyName postgres.ColumnString
	FactorType   postgres.ColumnString
	Status       postgres.ColumnString
	CreatedAt    postgres.ColumnTimestampz
	UpdatedAt    postgres.ColumnTimestampz
	Secret       postgres.ColumnString

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type MfaFactorsTable struct {
	mfaFactorsTable

	EXCLUDED mfaFactorsTable
}

// AS creates new MfaFactorsTable with assigned alias
func (a MfaFactorsTable) AS(alias string) *MfaFactorsTable {
	return newMfaFactorsTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new MfaFactorsTable with assigned schema name
func (a MfaFactorsTable) FromSchema(schemaName string) *MfaFactorsTable {
	return newMfaFactorsTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new MfaFactorsTable with assigned table prefix
func (a MfaFactorsTable) WithPrefix(prefix string) *MfaFactorsTable {
	return newMfaFactorsTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new MfaFactorsTable with assigned table suffix
func (a MfaFactorsTable) WithSuffix(suffix string) *MfaFactorsTable {
	return newMfaFactorsTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newMfaFactorsTable(schemaName, tableName, alias string) *MfaFactorsTable {
	return &MfaFactorsTable{
		mfaFactorsTable: newMfaFactorsTableImpl(schemaName, tableName, alias),
		EXCLUDED:        newMfaFactorsTableImpl("", "excluded", ""),
	}
}

func newMfaFactorsTableImpl(schemaName, tableName, alias string) mfaFactorsTable {
	var (
		IDColumn           = postgres.StringColumn("id")
		UserIDColumn       = postgres.StringColumn("user_id")
		FriendlyNameColumn = postgres.StringColumn("friendly_name")
		FactorTypeColumn   = postgres.StringColumn("factor_type")
		StatusColumn       = postgres.StringColumn("status")
		CreatedAtColumn    = postgres.TimestampzColumn("created_at")
		UpdatedAtColumn    = postgres.TimestampzColumn("updated_at")
		SecretColumn       = postgres.StringColumn("secret")
		allColumns         = postgres.ColumnList{IDColumn, UserIDColumn, FriendlyNameColumn, FactorTypeColumn, StatusColumn, CreatedAtColumn, UpdatedAtColumn, SecretColumn}
		mutableColumns     = postgres.ColumnList{UserIDColumn, FriendlyNameColumn, FactorTypeColumn, StatusColumn, CreatedAtColumn, UpdatedAtColumn, SecretColumn}
	)

	return mfaFactorsTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:           IDColumn,
		UserID:       UserIDColumn,
		FriendlyName: FriendlyNameColumn,
		FactorType:   FactorTypeColumn,
		Status:       StatusColumn,
		CreatedAt:    CreatedAtColumn,
		UpdatedAt:    UpdatedAtColumn,
		Secret:       SecretColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
