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

var RefreshTokens = newRefreshTokensTable("auth", "refresh_tokens", "")

type refreshTokensTable struct {
	postgres.Table

	// Columns
	InstanceID postgres.ColumnString
	ID         postgres.ColumnInteger
	Token      postgres.ColumnString
	UserID     postgres.ColumnString
	Revoked    postgres.ColumnBool
	CreatedAt  postgres.ColumnTimestampz
	UpdatedAt  postgres.ColumnTimestampz
	Parent     postgres.ColumnString
	SessionID  postgres.ColumnString

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type RefreshTokensTable struct {
	refreshTokensTable

	EXCLUDED refreshTokensTable
}

// AS creates new RefreshTokensTable with assigned alias
func (a RefreshTokensTable) AS(alias string) *RefreshTokensTable {
	return newRefreshTokensTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new RefreshTokensTable with assigned schema name
func (a RefreshTokensTable) FromSchema(schemaName string) *RefreshTokensTable {
	return newRefreshTokensTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new RefreshTokensTable with assigned table prefix
func (a RefreshTokensTable) WithPrefix(prefix string) *RefreshTokensTable {
	return newRefreshTokensTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new RefreshTokensTable with assigned table suffix
func (a RefreshTokensTable) WithSuffix(suffix string) *RefreshTokensTable {
	return newRefreshTokensTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newRefreshTokensTable(schemaName, tableName, alias string) *RefreshTokensTable {
	return &RefreshTokensTable{
		refreshTokensTable: newRefreshTokensTableImpl(schemaName, tableName, alias),
		EXCLUDED:           newRefreshTokensTableImpl("", "excluded", ""),
	}
}

func newRefreshTokensTableImpl(schemaName, tableName, alias string) refreshTokensTable {
	var (
		InstanceIDColumn = postgres.StringColumn("instance_id")
		IDColumn         = postgres.IntegerColumn("id")
		TokenColumn      = postgres.StringColumn("token")
		UserIDColumn     = postgres.StringColumn("user_id")
		RevokedColumn    = postgres.BoolColumn("revoked")
		CreatedAtColumn  = postgres.TimestampzColumn("created_at")
		UpdatedAtColumn  = postgres.TimestampzColumn("updated_at")
		ParentColumn     = postgres.StringColumn("parent")
		SessionIDColumn  = postgres.StringColumn("session_id")
		allColumns       = postgres.ColumnList{InstanceIDColumn, IDColumn, TokenColumn, UserIDColumn, RevokedColumn, CreatedAtColumn, UpdatedAtColumn, ParentColumn, SessionIDColumn}
		mutableColumns   = postgres.ColumnList{InstanceIDColumn, TokenColumn, UserIDColumn, RevokedColumn, CreatedAtColumn, UpdatedAtColumn, ParentColumn, SessionIDColumn}
	)

	return refreshTokensTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		InstanceID: InstanceIDColumn,
		ID:         IDColumn,
		Token:      TokenColumn,
		UserID:     UserIDColumn,
		Revoked:    RevokedColumn,
		CreatedAt:  CreatedAtColumn,
		UpdatedAt:  UpdatedAtColumn,
		Parent:     ParentColumn,
		SessionID:  SessionIDColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
