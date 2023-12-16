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

var SamlRelayStates = newSamlRelayStatesTable("auth", "saml_relay_states", "")

type samlRelayStatesTable struct {
	postgres.Table

	// Columns
	ID            postgres.ColumnString
	SsoProviderID postgres.ColumnString
	RequestID     postgres.ColumnString
	ForEmail      postgres.ColumnString
	RedirectTo    postgres.ColumnString
	FromIPAddress postgres.ColumnString
	CreatedAt     postgres.ColumnTimestampz
	UpdatedAt     postgres.ColumnTimestampz
	FlowStateID   postgres.ColumnString

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type SamlRelayStatesTable struct {
	samlRelayStatesTable

	EXCLUDED samlRelayStatesTable
}

// AS creates new SamlRelayStatesTable with assigned alias
func (a SamlRelayStatesTable) AS(alias string) *SamlRelayStatesTable {
	return newSamlRelayStatesTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new SamlRelayStatesTable with assigned schema name
func (a SamlRelayStatesTable) FromSchema(schemaName string) *SamlRelayStatesTable {
	return newSamlRelayStatesTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new SamlRelayStatesTable with assigned table prefix
func (a SamlRelayStatesTable) WithPrefix(prefix string) *SamlRelayStatesTable {
	return newSamlRelayStatesTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new SamlRelayStatesTable with assigned table suffix
func (a SamlRelayStatesTable) WithSuffix(suffix string) *SamlRelayStatesTable {
	return newSamlRelayStatesTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newSamlRelayStatesTable(schemaName, tableName, alias string) *SamlRelayStatesTable {
	return &SamlRelayStatesTable{
		samlRelayStatesTable: newSamlRelayStatesTableImpl(schemaName, tableName, alias),
		EXCLUDED:             newSamlRelayStatesTableImpl("", "excluded", ""),
	}
}

func newSamlRelayStatesTableImpl(schemaName, tableName, alias string) samlRelayStatesTable {
	var (
		IDColumn            = postgres.StringColumn("id")
		SsoProviderIDColumn = postgres.StringColumn("sso_provider_id")
		RequestIDColumn     = postgres.StringColumn("request_id")
		ForEmailColumn      = postgres.StringColumn("for_email")
		RedirectToColumn    = postgres.StringColumn("redirect_to")
		FromIPAddressColumn = postgres.StringColumn("from_ip_address")
		CreatedAtColumn     = postgres.TimestampzColumn("created_at")
		UpdatedAtColumn     = postgres.TimestampzColumn("updated_at")
		FlowStateIDColumn   = postgres.StringColumn("flow_state_id")
		allColumns          = postgres.ColumnList{IDColumn, SsoProviderIDColumn, RequestIDColumn, ForEmailColumn, RedirectToColumn, FromIPAddressColumn, CreatedAtColumn, UpdatedAtColumn, FlowStateIDColumn}
		mutableColumns      = postgres.ColumnList{SsoProviderIDColumn, RequestIDColumn, ForEmailColumn, RedirectToColumn, FromIPAddressColumn, CreatedAtColumn, UpdatedAtColumn, FlowStateIDColumn}
	)

	return samlRelayStatesTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:            IDColumn,
		SsoProviderID: SsoProviderIDColumn,
		RequestID:     RequestIDColumn,
		ForEmail:      ForEmailColumn,
		RedirectTo:    RedirectToColumn,
		FromIPAddress: FromIPAddressColumn,
		CreatedAt:     CreatedAtColumn,
		UpdatedAt:     UpdatedAtColumn,
		FlowStateID:   FlowStateIDColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
