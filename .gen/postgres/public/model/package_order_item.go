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

type PackageOrderItem struct {
	ID             int32 `sql:"primary_key"`
	PackageOrderID int32
	PackageID      int32
	Created        time.Time
}
