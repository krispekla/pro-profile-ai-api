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

var SsoProviders = newSsoProvidersTable("auth", "sso_providers", "")

type ssoProvidersTable struct {
	postgres.Table

	// Columns
	ID         postgres.ColumnString
	ResourceID postgres.ColumnString
	CreatedAt  postgres.ColumnTimestampz
	UpdatedAt  postgres.ColumnTimestampz

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type SsoProvidersTable struct {
	ssoProvidersTable

	EXCLUDED ssoProvidersTable
}

// AS creates new SsoProvidersTable with assigned alias
func (a SsoProvidersTable) AS(alias string) *SsoProvidersTable {
	return newSsoProvidersTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new SsoProvidersTable with assigned schema name
func (a SsoProvidersTable) FromSchema(schemaName string) *SsoProvidersTable {
	return newSsoProvidersTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new SsoProvidersTable with assigned table prefix
func (a SsoProvidersTable) WithPrefix(prefix string) *SsoProvidersTable {
	return newSsoProvidersTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new SsoProvidersTable with assigned table suffix
func (a SsoProvidersTable) WithSuffix(suffix string) *SsoProvidersTable {
	return newSsoProvidersTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newSsoProvidersTable(schemaName, tableName, alias string) *SsoProvidersTable {
	return &SsoProvidersTable{
		ssoProvidersTable: newSsoProvidersTableImpl(schemaName, tableName, alias),
		EXCLUDED:          newSsoProvidersTableImpl("", "excluded", ""),
	}
}

func newSsoProvidersTableImpl(schemaName, tableName, alias string) ssoProvidersTable {
	var (
		IDColumn         = postgres.StringColumn("id")
		ResourceIDColumn = postgres.StringColumn("resource_id")
		CreatedAtColumn  = postgres.TimestampzColumn("created_at")
		UpdatedAtColumn  = postgres.TimestampzColumn("updated_at")
		allColumns       = postgres.ColumnList{IDColumn, ResourceIDColumn, CreatedAtColumn, UpdatedAtColumn}
		mutableColumns   = postgres.ColumnList{ResourceIDColumn, CreatedAtColumn, UpdatedAtColumn}
	)

	return ssoProvidersTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:         IDColumn,
		ResourceID: ResourceIDColumn,
		CreatedAt:  CreatedAtColumn,
		UpdatedAt:  UpdatedAtColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
