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

type GeneratedPackageImg struct {
	ID                 int32 `sql:"primary_key"`
	GeneratedPackageID int32
	ImgURL             string
	ModelID            int32
	Created            time.Time
}
