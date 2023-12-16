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

var Package = newPackageTable("public", "package", "")

type packageTable struct {
	postgres.Table

	// Columns
	ID          postgres.ColumnInteger
	Name        postgres.ColumnString
	Description postgres.ColumnString
	CoverImgURL postgres.ColumnString
	Created     postgres.ColumnTimestamp

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type PackageTable struct {
	packageTable

	EXCLUDED packageTable
}

// AS creates new PackageTable with assigned alias
func (a PackageTable) AS(alias string) *PackageTable {
	return newPackageTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new PackageTable with assigned schema name
func (a PackageTable) FromSchema(schemaName string) *PackageTable {
	return newPackageTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new PackageTable with assigned table prefix
func (a PackageTable) WithPrefix(prefix string) *PackageTable {
	return newPackageTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new PackageTable with assigned table suffix
func (a PackageTable) WithSuffix(suffix string) *PackageTable {
	return newPackageTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newPackageTable(schemaName, tableName, alias string) *PackageTable {
	return &PackageTable{
		packageTable: newPackageTableImpl(schemaName, tableName, alias),
		EXCLUDED:     newPackageTableImpl("", "excluded", ""),
	}
}

func newPackageTableImpl(schemaName, tableName, alias string) packageTable {
	var (
		IDColumn          = postgres.IntegerColumn("id")
		NameColumn        = postgres.StringColumn("name")
		DescriptionColumn = postgres.StringColumn("description")
		CoverImgURLColumn = postgres.StringColumn("cover_img_url")
		CreatedColumn     = postgres.TimestampColumn("created")
		allColumns        = postgres.ColumnList{IDColumn, NameColumn, DescriptionColumn, CoverImgURLColumn, CreatedColumn}
		mutableColumns    = postgres.ColumnList{NameColumn, DescriptionColumn, CoverImgURLColumn, CreatedColumn}
	)

	return packageTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:          IDColumn,
		Name:        NameColumn,
		Description: DescriptionColumn,
		CoverImgURL: CoverImgURLColumn,
		Created:     CreatedColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}