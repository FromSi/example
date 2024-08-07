package tools

import (
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/tyler-smith/go-bip32"
	"github.com/tyler-smith/go-bip39"
)

type AddressBTC struct {
	address  string
	mnemonic string
}

func NewAddressBTC() *AddressBTC {
	entropy, _ := bip39.NewEntropy(128)
	mnemonic, _ := bip39.NewMnemonic(entropy)

	return &AddressBTC{
		address:  GenerateAddressBTC(mnemonic),
		mnemonic: mnemonic,
	}
}

func (addressBTC AddressBTC) GetAddress() string {
	return addressBTC.address
}

func (addressBTC AddressBTC) GetMnemonic() string {
	return addressBTC.mnemonic
}

func GenerateAddressBTC(mnemonic string) string {
	seed := bip39.NewSeed(mnemonic, "")

	masterKey, _ := bip32.NewMasterKey(seed)
	privateKey, _ := masterKey.Serialize()

	witnessProg := btcutil.Hash160(privateKey)
	addressWitnessPubKeyHash, _ := btcutil.NewAddressWitnessPubKeyHash(witnessProg, &chaincfg.MainNetParams)

	return addressWitnessPubKeyHash.EncodeAddress()
}
