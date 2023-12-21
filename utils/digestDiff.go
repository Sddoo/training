package utils

import (
	"crypto/sha256"
	"fmt"
)

func digestDiff(sha1, sha2 [32]byte) int {
	count := 0
	for i := 0; i < 32; i++ {
		diff := sha1[i] ^ sha2[i]
		for j := 0; j < 8; j++ {
			if diff >> j & 1 == 1 {
				count++
			}
		}
	}
	return count
}
