package pkg

import (
	"github.com/Ze-Victor/search-zip-code/internal/util"
)

type CredentialsAuth struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r *CredentialsAuth) Validate() error {
	if r.Email == "" {
		return util.ErrParamIsRequired("email", "string")
	}

	if r.Password == "" {
		return util.ErrParamIsRequired("password", "string")
	}

	return nil
}
