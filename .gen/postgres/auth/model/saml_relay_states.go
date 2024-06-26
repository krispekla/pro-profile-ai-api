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

type SamlRelayStates struct {
	ID            uuid.UUID `sql:"primary_key"`
	SsoProviderID uuid.UUID
	RequestID     string
	ForEmail      *string
	RedirectTo    *string
	FromIPAddress *string
	CreatedAt     *time.Time
	UpdatedAt     *time.Time
	FlowStateID   *uuid.UUID
}
