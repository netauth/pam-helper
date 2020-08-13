package module

import (
	"context"
)

// An Authenticator has the mechanism used for actually authenticating
// an entity.
type Authenticator interface {
	AuthEntity(context.Context, string, string) error
	SetServiceName(string)
}
