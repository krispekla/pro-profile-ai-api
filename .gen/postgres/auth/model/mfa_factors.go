//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

import (
	"github.com/google/uuid"
	"time"
)

type MfaFactors struct {
	ID           uuid.UUID `sql:"primary_key"`
	UserID       uuid.UUID
	FriendlyName *string
	FactorType   FactorType
	Status       FactorStatus
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Secret       *string
}
