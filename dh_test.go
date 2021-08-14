package dh

import (
	"math/big"
	"testing"
)

func TestOne(t *testing.T) {
	priv, _ := new(big.Int).SetString("54f3e39e6757434bb6c6ebc1636fdaa628357d4edbadfefa2540cfcc8d48d916572ee01cf3cc044fdecaa3774b449b5c5d502aaea4f20aa4ed601dcbf594705fc039aec15b8e7d1469fb95ba12c9addb126902dd1c5de21b93838337703d6217", 16)
	peerPub, _ := new(big.Int).SetString("b8a7702e82494e1af01e9c292c5daa716df817ffe845f7bcb12d7b7a1409976d4ecc8d287654b1839515afc2e24d84cc2ecbbec69efd19456cb0e90b95986e35e9ad80d2a5bc9e03e161afd7de9e03efd7fc77b17fe572fe7d34797ca0ab4731", 16)
	peerPriv, _ := new(big.Int).SetString("72e9145ddae3590977dbdfdd8d4f81d4a7b6a8bdf7a6376b02e54f0e240120683e5c0ef949e4699568176d274d3e06cdf42e8ba20f2ddcd90b80750dec8c8741d9c36434a2db58eca1fa7df31991f9199913fb37d01fbb03b4d43d80f22a0066", 16)
	dh := New(firstOakley, priv)
	shared := dh.SharedKey(peerPub)
	expectedShared, _ := new(big.Int).SetString("00cbf618713fedd19945eddbcefe07de2cfd88370d3b1abbd8dff8dbe2151db3786aeef3c3f328a9db9a63e998659a78e904630ec0dbb1e078f838c6993b6cac868c6387c0b47af1de464b639948169bf709b0d4cd063947785d9e1e6f646239", 16)

	if (expectedShared.Cmp(shared)) != 0 {
		t.Errorf("secret = %s; want %s", shared, expectedShared)
	}

	// peer should get same key
	dhPeer := New(firstOakley, peerPriv)
	peerShared := dhPeer.SharedKey(dh.PubKey())
	if (expectedShared.Cmp(peerShared)) != 0 {
		t.Errorf("peer secret = %s; want %s", peerShared, expectedShared)
	}

}
