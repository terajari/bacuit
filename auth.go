package bacuit

import (
	"context"
	"errors"
	"regexp"
	"strings"
)

var (
	UsernameMinLength   = 2
	PasswordMinLength   = 6
	EmailRegexp         = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	ErrInvalidUsername  = errors.New("username must be at least 2 characters long")
	ErrInvalidPassword  = errors.New("password must be at least 6 characters long")
	ErrInvalidEmail     = errors.New("email is invalid")
	ErrPasswordMismatch = errors.New("password mismatch")
)

type AuthService interface {
	Register(ctx context.Context, input RegisterInput) (AuthResponse, error)
}

type AuthResponse struct {
	Token string
}

type RegisterInput struct {
	Email           string
	Username        string
	Password        string
	ConfirmPassword string
}

func (ri *RegisterInput) Sanitize() {
	ri.Email = strings.TrimSpace(ri.Email)
	ri.Email = strings.ToLower(ri.Email)

	ri.Username = strings.TrimSpace(ri.Username)
}

func (ri *RegisterInput) Validate() error {
	if len(ri.Username) < UsernameMinLength {
		return ErrInvalidUsername
	}

	if len(ri.Password) < PasswordMinLength {
		return ErrInvalidPassword
	}

	if !EmailRegexp.MatchString(ri.Email) {
		return ErrInvalidEmail
	}

	if ri.Password != ri.ConfirmPassword {
		return ErrPasswordMismatch
	}

	return nil
}
