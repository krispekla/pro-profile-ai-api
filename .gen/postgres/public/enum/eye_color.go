//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package enum

import "github.com/go-jet/jet/v2/postgres"

var EyeColor = &struct {
	Brown postgres.StringExpression
	Blue  postgres.StringExpression
	Green postgres.StringExpression
	Gray  postgres.StringExpression
}{
	Brown: postgres.NewEnumValue("brown"),
	Blue:  postgres.NewEnumValue("blue"),
	Green: postgres.NewEnumValue("green"),
	Gray:  postgres.NewEnumValue("gray"),
}
