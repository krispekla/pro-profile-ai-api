//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package enum

import "github.com/go-jet/jet/v2/postgres"

var CodeChallengeMethod = &struct {
	S256  postgres.StringExpression
	Plain postgres.StringExpression
}{
	S256:  postgres.NewEnumValue("s256"),
	Plain: postgres.NewEnumValue("plain"),
}
