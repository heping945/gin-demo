package req

import (
	"github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"
)

// Custom claims structure
type CustomClaims struct {
	UUID     uuid.UUID
	ID       uint
	Username string
	jwt.StandardClaims
}
