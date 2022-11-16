package token

import (
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/johncalvinroberts/cryp/internal/errors"
)

// Create a struct that will be encoded to a JWT.
// We add jwt.RegisteredClaims as an embedded type, to provide fields like expiry time
type Claims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

type TokenService struct {
	Secret string
	// storageService   *storage.StorageService
}

func (svc *TokenService) IssueJWT(email string) (string, error) {
	expirationTime := time.Now().Add(10 * time.Minute)
	// Create the JWT claims, which includes the username and expiry time
	claims := &Claims{
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, err := token.SignedString([]byte((svc.Secret)))
	// TODO: persist to storage
	return tokenString, err
}

func (svc *TokenService) VerifyTokenAndParseClaims(tkn string) (*Claims, error) {
	claims := &Claims{}
	parsedToken, err := jwt.ParseWithClaims(tkn, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte((svc.Secret)), nil
	})
	// println(err.Error())
	// error parsing JWT
	if err != nil {
		return nil, err
	}
	// invalid token
	if !parsedToken.Valid {
		return nil, errors.ErrInvalidToken
	}
	// TODO: check in storage to make sure token hasn't been evicted
	return claims, nil
}

func (svc *TokenService) EvictToken(tkn string) error {
	// TODO: remove token from storage
	return nil
}
