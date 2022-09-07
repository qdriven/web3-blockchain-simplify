package hdwallet

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts"
)

const mnemonic = "tag volcano eight thank tide danger coast health above argue embrace heavy"

var wallet *Wallet

func init() {
	w, err := NewFromMnemonic(mnemonic)
	if err != nil {
		panic(err)
	}
	wallet = w
}

func DefaultWallet() *Wallet {
	return wallet
}

func Drive(hdIndex int) (accounts.Account, error) {
	path := derivationPath(hdIndex)
	return wallet.Derive(path, false)
}

// DerivationPath represents the computer friendly version of a hierarchical
// deterministic wallet account derivaion path.
//
// The BIP-32 spec https://github.com/bitcoin/bips/blob/master/bip-0032.mediawiki
// defines derivation paths to be of the form:
//
//   m / purpose' / coin_type' / account' / change / address_index
//
// The BIP-44 spec https://github.com/bitcoin/bips/blob/master/bip-0044.mediawiki
// defines that the `purpose` be 44' (or 0x8000002C) for crypto currencies, and
// SLIP-44 https://github.com/satoshilabs/slips/blob/master/slip-0044.md assigns
// the `coin_type` 60' (or 0x8000003C) to Ethereum.
//
// The root path for Ethereum is m/44'/60'/0'/0 according to the specification
// from https://github.com/ethereum/EIPs/issues/84, albeit it's not set in stone
// yet whether accounts should increment the last component or the children of
// that. We will go with the simpler approach of incrementing the last component.
func derivationPath(idx int) accounts.DerivationPath {
	format := fmt.Sprintf("m/44'/60'/0'/0/%d", idx)
	return MustParseDerivationPath(format)
}
