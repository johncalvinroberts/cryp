package config

import (
	"fmt"
	"log"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/joeshaw/envdecode"
	"github.com/joho/godotenv"
)

type AppConfig struct {
	Debug              bool     `env:"DEBUG,required"`
	GinMode            string   `env:"GIN_MODE,required"`
	JWTSecret          string   `env:"JWT_SECRET,required"`
	JWTTokenTTL        int      `env:"JWT_TOKEN_TTL,default=10"`
	EmailMaskSecret    string   `env:"EMAIL_MASK_SECRET,required"` // NOTE: This MUST be of 32 byte length
	Port               string   `env:"PORT,default=9000"`
	Timeout            int      `env:"TIMEOUT,default=8000"`
	EmailTransportName string   `env:"EMAIL_TRANSPORT_NAME,default=fs"`
	AllowOrigins       []string `env:"CORS_ALLOW_ORIGINS,required"`
	AWSSession         *session.Session
	AWS                struct {
		ID       string `env:"AWS_ACCESS_KEY_ID"`
		Secret   string `env:"AWS_SECRET_ACCESS_KEY,required"`
		Region   string `env:"AWS_REGION,required"`
		Endpoint string `env:"AWS_ENDPOINT,required"`
		Token    string `env:"AWS_TOKEN"`
	}
	Storage struct {
		WhoamiBucketName      string `env:"WHOAMI_BUCKET_NAME,required"`
		BlobBucketName        string `env:"BLOB_BUCKET_NAME,required"`
		BlobPointerBucketName string `env:"BLOB_POINTERS_BUCKET_NAME,required"`
	}
}

func InitAppConfig() *AppConfig {
	var c AppConfig
	godotenv.Load()

	if err := envdecode.StrictDecode(&c); err != nil {
		log.Fatalf("Failed to decode: %s", err)
	}

	massagedAllowOrigins := []string{}
	// godotenv is not properly serializing to slice, so need to manually massage
	for _, chunk := range c.AllowOrigins {
		split := strings.Split(chunk, ",")
		for _, str := range split {
			fmt.Println(str)
			massagedAllowOrigins = append(massagedAllowOrigins, strings.TrimSpace(str))
		}
	}
	c.AllowOrigins = massagedAllowOrigins
	c.AWSSession = session.Must(session.NewSession(&aws.Config{
		S3ForcePathStyle: aws.Bool(true),
		Region:           aws.String(c.AWS.Region),
		Credentials:      credentials.NewStaticCredentials(c.AWS.ID, c.AWS.Secret, c.AWS.Token),
		Endpoint:         aws.String(c.AWS.Endpoint),
	}))
	return &c
}
