package diffiehellman

import (
	"crypto/rand"
	"math/big"
)

// Diffie-Hellman-Merkle key exchange
// Private keys should be generated randomly.

func PrivateKey(p *big.Int) *big.Int {
	min := big.NewInt(2)
	max := big.NewInt(0).Sub(p, min)

	result, _ := rand.Int(rand.Reader, max)

	return result.Add(result, min)
}

func PublicKey(private, p *big.Int, g int64) *big.Int {
	bigG := big.NewInt(g)

	return big.NewInt(0).Exp(bigG, private, p)
}

func NewPair(p *big.Int, g int64) (*big.Int, *big.Int) {
	privateKey := PrivateKey(p)
	publicKey := PublicKey(privateKey, p, g)

	return privateKey, publicKey
}

func SecretKey(private1, public2, p *big.Int) *big.Int {
	return big.NewInt(0).Exp(public2, private1, p)
}
