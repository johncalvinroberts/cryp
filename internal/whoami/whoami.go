package whoami

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/johncalvinroberts/cryp/internal/email"
	"github.com/johncalvinroberts/cryp/internal/errors"
	"github.com/johncalvinroberts/cryp/internal/storage"
	"github.com/johncalvinroberts/cryp/internal/utils"
	"golang.org/x/net/context"
)

const (
	OTP_LENGTH              = 6
	WHOAMI_CHALLENGE_PREFIX = "whoami-challenge"
)

// Create a struct that will be encoded to a JWT.
// We add jwt.RegisteredClaims as an embedded type, to provide fields like expiry time
type Claims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

type WhoamiService struct {
	secret           string
	whoamiBucketName string
	storageService   *storage.StorageService
	emailService     *email.EmailService
}

func (svc *WhoamiService) FindWhoamiChallenge(email string) (string, error) {
	key := storage.ComposeKey(WHOAMI_CHALLENGE_PREFIX, email)
	ctx := context.Background()
	return svc.storageService.ReadToString(ctx, svc.whoamiBucketName, key)
}

func (svc *WhoamiService) issueJWT(email string) (string, error) {
	expirationTime := time.Now().Add(5 * time.Minute)
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
	tokenString, err := token.SignedString([]byte((svc.secret)))
	// TODO: persist to storage or redis
	return tokenString, err
}

// create otp + start whoami flow
func (svc *WhoamiService) StartWhoamiChallenge(email string) error {
	otp := utils.RandomSecret(OTP_LENGTH)
	ctx := context.Background()
	key := storage.ComposeKey(WHOAMI_CHALLENGE_PREFIX, email)
	_, err := svc.storageService.Write(ctx, svc.whoamiBucketName, key, strings.NewReader(otp))
	if err != nil {
		return err
	}
	msg := fmt.Sprintf("Your one-time password for Cryp: <code>%s</code>", otp)
	err = svc.emailService.SendANiceEmail(email, msg, "Cryp One-time password")
	if err != nil {
		return err
	}
	return nil
}

// redeem otp
func (svc *WhoamiService) TryWhoamiChallenge(email string, otp string) (string, error) {
	challenge, err := svc.FindWhoamiChallenge(email)
	if err != nil {
		log.Printf("encountered error when finding whoami challenge: %v\n", err)
		return "", errors.ErrInternalServerError
	}
	if otp != challenge || challenge == "" {
		log.Println("whoami challenge failed or not found")
		return "", errors.ErrWhoamiChallengeNotFound
	}
	// TODO: check expiration of whoami challenge
	jwt, err := svc.issueJWT(email)
	if err != nil {
		log.Printf("failed to issue jwt: %v\n", err)
		return "", errors.ErrInternalServerError
	}
	// cleanup
	defer svc.DestroyWhoamiChallenge(email)
	return jwt, nil
}

func (svc *WhoamiService) DestroyWhoamiChallenge(email string) {
	ctx := context.Background()
	key := storage.ComposeKey(WHOAMI_CHALLENGE_PREFIX, email)
	err := svc.storageService.Delete(ctx, svc.whoamiBucketName, key)
	if err != nil {
		log.Printf("failed to delete whoami challenge, key: %s", key)
	}
}

func (svc *WhoamiService) VerifyWhoami(endpointHandler func(c *gin.Context)) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenHeader := c.Request.Header.Get("Authorization")
		// no token in header
		if tokenHeader == "" {
			utils.RespondUnauthorized(c, errors.ErrUnauthorized)
			return
		}
		claims := &Claims{}

		token, err := jwt.ParseWithClaims(tokenHeader, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte((svc.secret)), nil
		})
		// error parsing JWT
		if err != nil {
			println(err.Error(), tokenHeader)
			utils.RespondUnauthorized(c, err)
			return
		}
		// invalid token
		if !token.Valid {
			utils.RespondUnauthorized(c, errors.ErrInvalidToken)
			return
		}
		endpointHandler(c)
	}
}

func InitWhoamiService(JWTSecret string, whoamiBucketName string, storageService *storage.StorageService, emailService *email.EmailService) *WhoamiService {
	return &WhoamiService{secret: JWTSecret, storageService: storageService, whoamiBucketName: whoamiBucketName, emailService: emailService}
}
