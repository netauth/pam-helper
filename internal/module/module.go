// Package module implements all the internals of setting up the
// request for authentication request and returning the result back to
// the call stack.
package module

import (
	"context"
	"errors"
	"io/ioutil"
	"os"
	"strings"

	"github.com/hashicorp/go-hclog"
)

// req contains all the metadata associated with a particular request.
// This is obtained from the environment, and from stdin.
type req struct {
	rType   string
	entity  string
	secret  string
	service string
}

func reqFromEnvironment() (*req, error) {
	r := new(req)

	switch os.Getenv("PAM_TYPE") {
	case "account":
		r.rType = "account"
	case "auth":
		r.rType = "auth"
	default:
		return nil, errors.New("only account and auth types are supported")
	}

	r.entity = os.Getenv("PAM_USER")
	r.service = os.Getenv("PAM_SERVICE")

	return r, nil
}

func (r *req) getSecret() error {
	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		return err
	}
	r.secret = strings.TrimSuffix(string(b), string('\x00'))

	return nil
}

// Exec is the entrypoint to the module's code.  It is called directly
// by main.
func Exec(l hclog.Logger, a Authenticator) int {
	r, err := reqFromEnvironment()
	if err != nil {
		l.Debug("Error constructing request", "error", err)
		return 1
	}
	if err := r.getSecret(); err != nil {
		l.Debug("Error reading secret", "error", err)
		return 2
	}

	a.SetServiceName(r.service)
	if err := a.AuthEntity(context.Background(), r.entity, r.secret); err != nil {
		l.Debug("Authentication failed", "error", err)
		return 1
	}

	return 0
}
