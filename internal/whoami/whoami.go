package whoami

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
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

type WhoamiService struct {
	secret           string
	whoamiBucketName string
	storageService   *storage.StorageService
	emailService     *email.EmailService
}

func (svc *WhoamiService) FindWhoamiChallenge(email string) *string {
	var challenge string
	// TODO: lookup from storage
	return &challenge
}

func (svc *WhoamiService) issueJWT(email string) (string, error) {
	token := jwt.New(jwt.SigningMethodEdDSA)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(10 * time.Minute)
	claims["authorized"] = true
	claims["email"] = email
	tokenString, err := token.SignedString(svc.secret)
	if err != nil {
		return "Signing Error", err
	}
	// TODO: persist to storage or redis
	return tokenString, nil
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
	challenge := svc.FindWhoamiChallenge(email)
	if otp != *challenge || challenge == nil {
		log.Default().Print("otp challenge failed or not found")
		return "", errors.ErrValidationFailure
	}
	// TODO: check expiration of whoami challenge
	// cleanup
	defer svc.DestroyWhoamiChallenge(email)
	jwt, err := svc.issueJWT(email)
	if err != nil {
		return "", err
	}
	return jwt, nil
}

func (svc *WhoamiService) DestroyWhoamiChallenge(email string) {
	ctx := context.Background()
	key := storage.ComposeKey(WHOAMI_CHALLENGE_PREFIX, email)
	svc.storageService.Delete(ctx, svc.whoamiBucketName, key)
}

func InitWhoamiService(JWTSecret string, whoamiBucketName string, storageService *storage.StorageService, emailService *email.EmailService) *WhoamiService {
	return &WhoamiService{secret: JWTSecret, storageService: storageService, whoamiBucketName: whoamiBucketName, emailService: emailService}
}
