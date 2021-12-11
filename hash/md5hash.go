package hash

import (
	"crypto/md5"
	"math/big"
)

const OutputSpace = 100

func hash(key string) *big.Int {
	bi := big.NewInt(0)
	h := md5.New()
	h.Write([]byte(key))
	hash := bi.SetBytes(h.Sum(nil))
	return hash
}
func mod(num *big.Int) int {
	res := big.NewInt(0)
	mod := big.NewInt(OutputSpace)
	res.Mod(num, mod)
	return int(res.Int64())
}
func LocationOnRing(key string) int {
	hash := hash(key)
	return mod(hash)
}
