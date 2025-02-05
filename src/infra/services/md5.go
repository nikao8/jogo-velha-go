package infra_services

import (
	"crypto/md5"
	"encoding/hex"
)

type md5Service struct{}

func NewMD5Service() *md5Service {
	return &md5Service{}
}

func (s md5Service) Hash(value string) string {
	hash := md5.Sum([]byte(value))
	return hex.EncodeToString(hash[:])
}
