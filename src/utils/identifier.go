package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func MakePassword(rawPassword string) (string, error) {
	return MakeMD5(rawPassword), nil
}

func CheckPassword(rawPassword, encryptPassword string) (bool, error) {
	return MakeMD5(rawPassword) == encryptPassword, nil
}

func MakeMD5(raw string) string {
	hash := md5.New()
	hash.Write([]byte(raw))
	return hex.EncodeToString(hash.Sum(nil))
}
