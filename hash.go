package libada

import (
	"golang.org/x/crypto/blake2b"
	"golang.org/x/crypto/sha3"
)

func Blake2b224(raw []byte) []byte {
	hasher, _ := blake2b.New(28, nil)
	_, _ = hasher.Write(raw)
	return hasher.Sum(nil)
}

func Sha3AndBlake2b224(raw []byte) []byte {
	res := sha3.Sum256(raw)
	return Blake2b224(res[:])
}
