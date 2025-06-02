package aoc

import (
	"crypto/md5"
	"fmt"
)

func MD5Hash(word string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(word)))
}
