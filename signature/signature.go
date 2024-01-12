package signature

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/zeebo/blake3"
)

func GenerateKey() (*big.Int, *big.Int, *big.Int) {
	// Generate random big number prime for p
	p, err := rand.Prime(rand.Reader, 1024)
	if err != nil {
		fmt.Println(err.Error())
	}

	// Generate random big number prime for q
	q, err := rand.Prime(rand.Reader, 1024)
	if err != nil {
		fmt.Println(err.Error())
	}

	// Generate kunci e
	e, err := rand.Prime(rand.Reader, 512)
	if err != nil {
		fmt.Println(err.Error())
	}

	// n = p * q
	n := new(big.Int).Mul(p, q)

	// totient(n) = (p - 1) * (q - 1)
	pMin1 := new(big.Int).Sub(p, new(big.Int).SetInt64(1))
	qMin1 := new(big.Int).Sub(q, new(big.Int).SetInt64(1))
	totient := new(big.Int).Mul(pMin1, qMin1)

	// kunci d
	d := new(big.Int).ModInverse(e, totient)

	return e, d, n
}

func HashMsg(message []byte) *big.Int {
	msgDigest := blake3.Sum256([]byte(message))

	strMsgDigest := hex.EncodeToString(msgDigest[:])

	decimalmsgDigest := new(big.Int)
	decimalmsgDigest.SetString(strMsgDigest, 16)

	return decimalmsgDigest
}

func GenerateSignature(MD, d, n *big.Int) string {
	// Enkripsi = (md ** d) mod n
	signature := new(big.Int).Exp(MD, d, n)

	return fmt.Sprintf("%x", signature)
}

func DecryptSignature(signature string, e, n *big.Int) string {
	// Dekripsi = (signature ** e) mod n
	decimalSignature := new(big.Int)
	decimalSignature.SetString(signature, 16)

	msgDigest := new(big.Int).Exp(decimalSignature, e, n)

	return fmt.Sprintf("%x", msgDigest)
}
