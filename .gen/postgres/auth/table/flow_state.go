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

var FlowState = newFlowStateTable("auth", "flow_state", "")

type flowStateTable struct {
	postgres.Table

	// Columns
	ID                   postgres.ColumnString
	UserID               postgres.ColumnString
	AuthCode             postgres.ColumnString
	CodeChallengeMethod  postgres.ColumnString
	CodeChallenge        postgres.ColumnString
	ProviderType         postgres.ColumnString
	ProviderAccessToken  postgres.ColumnString
	ProviderRefreshToken postgres.ColumnString
	CreatedAt            postgres.ColumnTimestampz
	UpdatedAt            postgres.ColumnTimestampz
	AuthenticationMethod postgres.ColumnString

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type FlowStateTable struct {
	flowStateTable

	EXCLUDED flowStateTable
}

// AS creates new FlowStateTable with assigned alias
func (a FlowStateTable) AS(alias string) *FlowStateTable {
	return newFlowStateTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new FlowStateTable with assigned schema name
func (a FlowStateTable) FromSchema(schemaName string) *FlowStateTable {
	return newFlowStateTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new FlowStateTable with assigned table prefix
func (a FlowStateTable) WithPrefix(prefix string) *FlowStateTable {
	return newFlowStateTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new FlowStateTable with assigned table suffix
func (a FlowStateTable) WithSuffix(suffix string) *FlowStateTable {
	return newFlowStateTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newFlowStateTable(schemaName, tableName, alias string) *FlowStateTable {
	return &FlowStateTable{
		flowStateTable: newFlowStateTableImpl(schemaName, tableName, alias),
		EXCLUDED:       newFlowStateTableImpl("", "excluded", ""),
	}
}

func newFlowStateTableImpl(schemaName, tableName, alias string) flowStateTable {
	var (
		IDColumn                   = postgres.StringColumn("id")
		UserIDColumn               = postgres.StringColumn("user_id")
		AuthCodeColumn             = postgres.StringColumn("auth_code")
		CodeChallengeMethodColumn  = postgres.StringColumn("code_challenge_method")
		CodeChallengeColumn        = postgres.StringColumn("code_challenge")
		ProviderTypeColumn         = postgres.StringColumn("provider_type")
		ProviderAccessTokenColumn  = postgres.StringColumn("provider_access_token")
		ProviderRefreshTokenColumn = postgres.StringColumn("provider_refresh_token")
		CreatedAtColumn            = postgres.TimestampzColumn("created_at")
		UpdatedAtColumn            = postgres.TimestampzColumn("updated_at")
		AuthenticationMethodColumn = postgres.StringColumn("authentication_method")
		allColumns                 = postgres.ColumnList{IDColumn, UserIDColumn, AuthCodeColumn, CodeChallengeMethodColumn, CodeChallengeColumn, ProviderTypeColumn, ProviderAccessTokenColumn, ProviderRefreshTokenColumn, CreatedAtColumn, UpdatedAtColumn, AuthenticationMethodColumn}
		mutableColumns             = postgres.ColumnList{UserIDColumn, AuthCodeColumn, CodeChallengeMethodColumn, CodeChallengeColumn, ProviderTypeColumn, ProviderAccessTokenColumn, ProviderRefreshTokenColumn, CreatedAtColumn, UpdatedAtColumn, AuthenticationMethodColumn}
	)

	return flowStateTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:                   IDColumn,
		UserID:               UserIDColumn,
		AuthCode:             AuthCodeColumn,
		CodeChallengeMethod:  CodeChallengeMethodColumn,
		CodeChallenge:        CodeChallengeColumn,
		ProviderType:         ProviderTypeColumn,
		ProviderAccessToken:  ProviderAccessTokenColumn,
		ProviderRefreshToken: ProviderRefreshTokenColumn,
		CreatedAt:            CreatedAtColumn,
		UpdatedAt:            UpdatedAtColumn,
		AuthenticationMethod: AuthenticationMethodColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
