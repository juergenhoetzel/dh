package dh

import (
	"math/big"
)

type dh struct {
	n       *big.Int
	private *big.Int
}

var g *big.Int

type dhGroup int

const (
	firstOakley dhGroup = iota
	secondOakley
)

var group map[dhGroup]*big.Int

func init() {
	g = new(big.Int).SetInt64(2)
	group = map[dhGroup]*big.Int{}
	group[firstOakley], _ = new(big.Int).SetString("FFFFFFFFFFFFFFFFC90FDAA22168C234C4C6628B80DC1CD129024E088A67CC74020BBEA63B139B22514A08798E3404DDEF9519B3CD3A431B302B0A6DF25F14374FE1356D6D51C245E485B576625E7EC6F44C42E9A63A3620FFFFFFFFFFFFFFFF", 16)
	group[secondOakley], _ = new(big.Int).SetString("FFFFFFFFFFFFFFFFC90FDAA22168C234C4C6628B80DC1CD129024E088A67CC74020BBEA63B139B22514A08798E3404DDEF9519B3CD3A431B302B0A6DF25F14374FE1356D6D51C245E485B576625E7EC6F44C42E9A637ED6B0BFF5CB6F406B7EDEE386BFB5A899FA5AE9F24117C4B1FE649286651ECE65381FFFFFFFFFFFFFFFF", 16)
}
func New(g dhGroup, private *big.Int) dh {
	return dh{n: group[g], private: private}
}

func (d *dh) PubKey() *big.Int {
	return new(big.Int).Exp(g, d.private, d.n)
}

// SharedKey return the shared secret key based on the peers public key p
func (d *dh) SharedKey(pub *big.Int) *big.Int {
	return new(big.Int).Exp(pub, d.private, d.n)
}
