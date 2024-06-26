//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

import "errors"

type FactorType string

const (
	FactorType_Totp     FactorType = "totp"
	FactorType_Webauthn FactorType = "webauthn"
)

func (e *FactorType) Scan(value interface{}) error {
	var enumValue string
	switch val := value.(type) {
	case string:
		enumValue = val
	case []byte:
		enumValue = string(val)
	default:
		return errors.New("jet: Invalid scan value for AllTypesEnum enum. Enum value has to be of type string or []byte")
	}

	switch enumValue {
	case "totp":
		*e = FactorType_Totp
	case "webauthn":
		*e = FactorType_Webauthn
	default:
		return errors.New("jet: Invalid scan value '" + enumValue + "' for FactorType enum")
	}

	return nil
}

func (e FactorType) String() string {
	return string(e)
}
