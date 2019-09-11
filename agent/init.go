// DO NOT EDIT -- generated code

// Package agent - the agent communicates with pinpoint cloud to send integration data
package agent

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

// Generate a random string of A-Z chars with len = l
func randomString(len int) string {
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		bytes[i] = byte(randomInt(65, 90))
	}
	return string(bytes)
}

var emptyString string
var emptyInt int64
var emptyBool bool
var emptyFloat float64

// ModelWithID is a model interface that has a GetID method
type ModelWithID interface {
	GetID() string
}
