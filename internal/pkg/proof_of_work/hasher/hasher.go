package hasher

import "crypto/sha256"

type Sha256Hasher struct {
}

func New() Sha256Hasher {
	return Sha256Hasher{}
}

func (h Sha256Hasher) Hash(data []byte) []byte {
	res := sha256.Sum256(data)
	return res[:]
}
