package config

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/joeshaw/envdecode"
	"github.com/joho/godotenv"
)

type AppConfig struct {
	Debug      bool   `env:"DEBUG,required"`
	GinMode    string `env:"GIN_MODE,required"`
	JWTSecret  string `env:"JWT_SECRET,required"`
	Port       string `env:"PORT,default=9000"`
	Timeout    int    `env:"TIMEOUT,default=8000"`
	AWSSession *session.Session
	AWS        struct {
		ID       string `env:"AWS_ACCESS_KEY_ID"`
		Secret   string `env:"AWS_SECRET_ACCESS_KEY,required"`
		Region   string `env:"AWS_REGION,required"`
		Endpoint string `env:"AWS_ENDPOINT,required"`
	}
	Storage struct {
		WhoamiBucketName string `env:"WHOAMI_BUCKET_NAME,required"`
	}
}

func InitAppConfig() *AppConfig {
	var c AppConfig
	godotenv.Load()
	if err := envdecode.StrictDecode(&c); err != nil {
		log.Fatalf("Failed to decode: %s", err)
	}

	c.AWSSession = session.Must(session.NewSession(&aws.Config{
		S3ForcePathStyle: aws.Bool(true),
		Region:           aws.String(c.AWS.Region),
		Credentials:      &credentials.Credentials{},
		Endpoint:         aws.String(c.AWS.Endpoint),
	}))
	return &c
}
