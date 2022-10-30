package email

import (
	"fmt"
	"os"
	"time"

	"github.com/johncalvinroberts/cryp/internal/config"
)

const (
	emailsDir = "emails"
)

type FSTransport struct{}

func (t *FSTransport) SendANiceEmail(to string, msg string, subject string, html string) error {
	key := time.Now().Unix()
	fileName := fmt.Sprintf("%s/%d", emailsDir, key)
	err := os.WriteFile(fileName, []byte(html), 0644)
	return err
}

func (t *FSTransport) InitTransport(_ *config.AppConfig) {
	// noop
}
