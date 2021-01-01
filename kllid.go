package kllid

import (
	"math/rand"
	"strings"
	"time"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func init() {
	rand.Seed(time.Now().UnixNano())
}
// New returns Random string of alpha numeric string
func New(idLen int) string {
	return randStringBytesRmndr(idLen)
}

func randStringBytesRmndr(idLen int) string {
	sb := strings.Builder{}
	sb.Grow(idLen)
	for i := 0; i < idLen; i++ {
		sb.WriteByte(letterBytes[rand.Int63()%int64(len(letterBytes))])
	}
	return sb.String()
}
