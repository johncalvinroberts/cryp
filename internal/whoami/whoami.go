package whoami

import (
	"fmt"
	"log"
	"strings"
	"time"

	"context"

	"github.com/gin-gonic/gin"
	"github.com/johncalvinroberts/cryp/internal/email"
	"github.com/johncalvinroberts/cryp/internal/errors"
	"github.com/johncalvinroberts/cryp/internal/storage"
	"github.com/johncalvinroberts/cryp/internal/token"
	"github.com/johncalvinroberts/cryp/internal/utils"
)

const (
	OTP_LENGTH              = 6
	WHOAMI_CHALLENGE_PREFIX = "whoami-challenge"
	CTX_JWT_KEY             = "CTX_JWT_KEY"
	CTX_JWT_CLAIMS_KEY      = "CTX_JWT_CLAIMS_KEY"
)

type WhoamiService struct {
	secret           string
	whoamiBucketName string
	storageService   *storage.StorageService
	emailService     *email.EmailService
	tokenService     *token.TokenService
}

func (svc *WhoamiService) FindWhoamiChallenge(email string) (string, error) {
	key := storage.ComposeKey(WHOAMI_CHALLENGE_PREFIX, email)
	ctx := context.Background()
	return svc.storageService.ReadToString(ctx, svc.whoamiBucketName, key)
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
		// cheap throttling
		time.Sleep(2 * time.Second)
		return "", errors.ErrWhoamiChallengeNotFound
	}
	// TODO: check expiration of whoami challenge
	jwt, err := svc.tokenService.IssueJWT(email)
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

func (svc *WhoamiService) VerifyWhoamiMiddleware(endpointHandler func(c *gin.Context)) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		// no token in header
		if token == "" {
			utils.RespondUnauthorized(c, errors.ErrUnauthorized)
			return
		}
		claims, err := svc.tokenService.VerifyTokenAndParseClaims(token)
		// error parsing JWT
		if err != nil {
			utils.RespondUnauthorized(c, err)
			return
		}
		c.Set(CTX_JWT_KEY, token)
		c.Set(CTX_JWT_CLAIMS_KEY, &claims)
		endpointHandler(c)
	}
}

func InitWhoamiService(JWTSecret string, whoamiBucketName string, storageService *storage.StorageService, emailService *email.EmailService) *WhoamiService {
	return &WhoamiService{
		secret:           JWTSecret,
		storageService:   storageService,
		whoamiBucketName: whoamiBucketName,
		emailService:     emailService,
		tokenService: &token.TokenService{
			Secret: JWTSecret,
		},
	}
}
