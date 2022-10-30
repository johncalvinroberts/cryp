package email

import (
	"fmt"

	"github.com/johncalvinroberts/cryp/internal/config"
)

const (
	CharSet = "UTF-8"
	From    = "no-reply@cryp.sh"
)

type EmailTransport interface {
	InitTransport(config *config.AppConfig)
	SendANiceEmail(to string, msg string, subject string, html string) error
}

type EmailService struct {
	transport EmailTransport
}

func (svc *EmailService) SendANiceEmail(to string, msg string, subject string) error {
	html := svc.BuildHtml(msg)
	return svc.transport.SendANiceEmail(to, msg, subject, html)
}

func (svc *EmailService) BuildHtml(msg string) string {
	return fmt.Sprintf(`
  <div className="email" style="
    border: 1px solid black;
    padding: 20px;
    font-family: sans-serif;
    line-height: 2;
    font-size: 20px;
  ">
    <h2>Hello There!</h2>
    <p>%s</p>
    <hr/>
  </div>
	`, msg)
}

func InitEmailService(config *config.AppConfig) *EmailService {
	var transport EmailTransport
	if config.EmailTransportName == "ses" {
		transport = &SESTRansport{}
	}
	if config.EmailTransportName == "fs" {
		transport = &FSTransport{}
	}
	transport.InitTransport(config)
	return &EmailService{transport}
}
