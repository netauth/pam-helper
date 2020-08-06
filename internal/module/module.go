// Package module implements all the internals of setting up the
// request for authentication request and returning the result back to
// the call stack.
package module

import (
	"bufio"
	"errors"
	"log"
	"os"
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
	rdr := bufio.NewReader(os.Stdin)

	b, _, err := rdr.ReadLine()
	if err != nil {
		return err
	}
	r.secret = string(b[:])

	return nil
}

// Exec is the entrypoint to the module's code.  It is called directly
// by main.
func Exec() int {
	r, err := reqFromEnvironment()
	if err != nil {
		log.Println(err)
		return 1
	}
	if err := r.getSecret(); err != nil {
		log.Println(err)
		return 2
	}

	log.Printf("%+v", r)

	return 0
}
