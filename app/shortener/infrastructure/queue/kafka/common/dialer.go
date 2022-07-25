package common

import (
	"crypto/tls"
	"crypto/x509"
	"github.com/creepzed/url-shortener-service/app/shared/infrastructure/log"
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/plain"
)

func GetDialer(username string, password string) *kafka.Dialer {
	if len(username) == 0 || len(password) == 0 {
		return kafka.DefaultDialer
	}

	validUserNameAndPassword(username, password)

	rootCAs, _ := x509.SystemCertPool()
	if rootCAs == nil {
		rootCAs = x509.NewCertPool()
	}
	dialer := &kafka.Dialer{

		DualStack: true,
		SASLMechanism: plain.Mechanism{
			Username: username, // access key
			Password: password, // secret
		},
		TLS: &tls.Config{
			InsecureSkipVerify: true,
			RootCAs:            rootCAs,
		},
	}
	return dialer
}

func validUserNameAndPassword(kafkaUsername string, kafkaPassword string) {
	if len(kafkaUsername) == 0 || len(kafkaPassword) == 0 {
		log.Fatal("username and password are required to connect to kafka")
	}
}
