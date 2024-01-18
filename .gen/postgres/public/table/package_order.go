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

var PackageOrder = newPackageOrderTable("public", "package_order", "")

type packageOrderTable struct {
	postgres.Table

	// Columns
	ID                      postgres.ColumnInteger
	UserID                  postgres.ColumnString
	Created                 postgres.ColumnTimestamp
	TotalAmount             postgres.ColumnInteger
	Currency                postgres.ColumnString
	Status                  postgres.ColumnString
	CouponID                postgres.ColumnInteger
	StripePaymentIntentID   postgres.ColumnString
	StripeCheckoutSessionID postgres.ColumnString

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type PackageOrderTable struct {
	packageOrderTable

	EXCLUDED packageOrderTable
}

// AS creates new PackageOrderTable with assigned alias
func (a PackageOrderTable) AS(alias string) *PackageOrderTable {
	return newPackageOrderTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new PackageOrderTable with assigned schema name
func (a PackageOrderTable) FromSchema(schemaName string) *PackageOrderTable {
	return newPackageOrderTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new PackageOrderTable with assigned table prefix
func (a PackageOrderTable) WithPrefix(prefix string) *PackageOrderTable {
	return newPackageOrderTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new PackageOrderTable with assigned table suffix
func (a PackageOrderTable) WithSuffix(suffix string) *PackageOrderTable {
	return newPackageOrderTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newPackageOrderTable(schemaName, tableName, alias string) *PackageOrderTable {
	return &PackageOrderTable{
		packageOrderTable: newPackageOrderTableImpl(schemaName, tableName, alias),
		EXCLUDED:          newPackageOrderTableImpl("", "excluded", ""),
	}
}

func newPackageOrderTableImpl(schemaName, tableName, alias string) packageOrderTable {
	var (
		IDColumn                      = postgres.IntegerColumn("id")
		UserIDColumn                  = postgres.StringColumn("user_id")
		CreatedColumn                 = postgres.TimestampColumn("created")
		TotalAmountColumn             = postgres.IntegerColumn("total_amount")
		CurrencyColumn                = postgres.StringColumn("currency")
		StatusColumn                  = postgres.StringColumn("status")
		CouponIDColumn                = postgres.IntegerColumn("coupon_id")
		StripePaymentIntentIDColumn   = postgres.StringColumn("stripe_payment_intent_id")
		StripeCheckoutSessionIDColumn = postgres.StringColumn("stripe_checkout_session_id")
		allColumns                    = postgres.ColumnList{IDColumn, UserIDColumn, CreatedColumn, TotalAmountColumn, CurrencyColumn, StatusColumn, CouponIDColumn, StripePaymentIntentIDColumn, StripeCheckoutSessionIDColumn}
		mutableColumns                = postgres.ColumnList{UserIDColumn, CreatedColumn, TotalAmountColumn, CurrencyColumn, StatusColumn, CouponIDColumn, StripePaymentIntentIDColumn, StripeCheckoutSessionIDColumn}
	)

	return packageOrderTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:                      IDColumn,
		UserID:                  UserIDColumn,
		Created:                 CreatedColumn,
		TotalAmount:             TotalAmountColumn,
		Currency:                CurrencyColumn,
		Status:                  StatusColumn,
		CouponID:                CouponIDColumn,
		StripePaymentIntentID:   StripePaymentIntentIDColumn,
		StripeCheckoutSessionID: StripeCheckoutSessionIDColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
