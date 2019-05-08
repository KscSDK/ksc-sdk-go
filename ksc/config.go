package ksc

import (
	"github.com/aws/aws-sdk-go/aws/credentials"
)

type Config struct {
	Credentials *credentials.Credentials

	Region *string
}
