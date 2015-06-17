package mturk

import (
	"gopkg.in/bulletind/goamz.v1/aws"
)

func Sign(auth aws.Auth, service, method, timestamp string, params map[string]string) {
	sign(auth, service, method, timestamp, params)
}
