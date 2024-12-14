package rkey

import (
	"math/big"

	_ "golang.org/x/crypto/ripemd160" // Required to use SumSha256Ripemd160
	"hg.sr.ht/~dchapes/ripple/crypto/doublehash"
)

// An AccountId is a hash of a public key.
//
// Its Ripple base58 encoding starts with `r` as is the Ripple address.
//
// TODO(dchapes): remove?
type AccountId struct {
	Id *big.Int
}

func (a *AccountId) raw() []byte { return a.Id.Bytes() }
func (a *AccountId) setraw(raw []byte) error {
	if a.Id == nil {
		a.Id = new(big.Int)
	}
	a.Id.SetBytes(raw)
	return nil
}

func accountIdHash(b []byte) *AccountId {
	sum := doublehash.SumSha256Ripemd160(b)
	return &AccountId{new(big.Int).SetBytes(sum[:])}
}