package domain

import (
	"strings"
)

type User struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (u *User) Validate() error {
	if strings.TrimSpace(u.Name) == "" {
		return ErrNameRequired
	}

	if !strings.Contains(u.Email, "@") {
		return ErrInvalidEmail
	}

	return nil
}
