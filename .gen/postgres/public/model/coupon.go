//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

import (
	"time"
)

type Coupon struct {
	ID         int32 `sql:"primary_key"`
	PackageID  int32
	Code       string
	Amount     *int32
	Currency   *string
	Percentage *int32
	Created    time.Time
}
