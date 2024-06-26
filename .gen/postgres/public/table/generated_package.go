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

var GeneratedPackage = newGeneratedPackageTable("public", "generated_package", "")

type generatedPackageTable struct {
	postgres.Table

	// Columns
	ID                 postgres.ColumnInteger
	PackageOrderItemID postgres.ColumnInteger
	CharacterID        postgres.ColumnInteger
	Status             postgres.ColumnString
	CoverImgURL        postgres.ColumnString
	Created            postgres.ColumnTimestamp
	Updated            postgres.ColumnTimestamp

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type GeneratedPackageTable struct {
	generatedPackageTable

	EXCLUDED generatedPackageTable
}

// AS creates new GeneratedPackageTable with assigned alias
func (a GeneratedPackageTable) AS(alias string) *GeneratedPackageTable {
	return newGeneratedPackageTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new GeneratedPackageTable with assigned schema name
func (a GeneratedPackageTable) FromSchema(schemaName string) *GeneratedPackageTable {
	return newGeneratedPackageTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new GeneratedPackageTable with assigned table prefix
func (a GeneratedPackageTable) WithPrefix(prefix string) *GeneratedPackageTable {
	return newGeneratedPackageTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new GeneratedPackageTable with assigned table suffix
func (a GeneratedPackageTable) WithSuffix(suffix string) *GeneratedPackageTable {
	return newGeneratedPackageTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newGeneratedPackageTable(schemaName, tableName, alias string) *GeneratedPackageTable {
	return &GeneratedPackageTable{
		generatedPackageTable: newGeneratedPackageTableImpl(schemaName, tableName, alias),
		EXCLUDED:              newGeneratedPackageTableImpl("", "excluded", ""),
	}
}

func newGeneratedPackageTableImpl(schemaName, tableName, alias string) generatedPackageTable {
	var (
		IDColumn                 = postgres.IntegerColumn("id")
		PackageOrderItemIDColumn = postgres.IntegerColumn("package_order_item_id")
		CharacterIDColumn        = postgres.IntegerColumn("character_id")
		StatusColumn             = postgres.StringColumn("status")
		CoverImgURLColumn        = postgres.StringColumn("cover_img_url")
		CreatedColumn            = postgres.TimestampColumn("created")
		UpdatedColumn            = postgres.TimestampColumn("updated")
		allColumns               = postgres.ColumnList{IDColumn, PackageOrderItemIDColumn, CharacterIDColumn, StatusColumn, CoverImgURLColumn, CreatedColumn, UpdatedColumn}
		mutableColumns           = postgres.ColumnList{PackageOrderItemIDColumn, CharacterIDColumn, StatusColumn, CoverImgURLColumn, CreatedColumn, UpdatedColumn}
	)

	return generatedPackageTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:                 IDColumn,
		PackageOrderItemID: PackageOrderItemIDColumn,
		CharacterID:        CharacterIDColumn,
		Status:             StatusColumn,
		CoverImgURL:        CoverImgURLColumn,
		Created:            CreatedColumn,
		Updated:            UpdatedColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
