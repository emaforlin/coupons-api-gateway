package entities

import (
	"errors"
	"slices"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

type CustomClaims struct {
	Role string `json:"role"`
	jwt.RegisteredClaims
}

// Valid implements jwt.Claims.
func (c *CustomClaims) Valid() error {
	var validRoles = []string{"customer", "foodplace", "visitor"}
	if slices.Index(validRoles, strings.ToLower(c.Role)) == -1 {
		return errors.New("invalid role")
	}
	return nil
}
