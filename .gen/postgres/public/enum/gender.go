//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package enum

import "github.com/go-jet/jet/v2/postgres"

var Gender = &struct {
	Male   postgres.StringExpression
	Female postgres.StringExpression
	Other  postgres.StringExpression
}{
	Male:   postgres.NewEnumValue("male"),
	Female: postgres.NewEnumValue("female"),
	Other:  postgres.NewEnumValue("other"),
}
