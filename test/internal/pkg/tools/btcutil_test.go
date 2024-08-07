package tools_test

import (
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
	. "github.com/fromsi/example/internal/pkg/tools"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Address BTC", func() {
	It("can make correctly address", func() {
		address := NewAddressBTC()

		_, err := btcutil.DecodeAddress(address.GetAddress(), &chaincfg.MainNetParams)

		Expect(err).To(BeNil())
	})

	It("can make correctly mnemonic", func() {
		address := NewAddressBTC()

		Expect(GenerateAddressBTC(address.GetMnemonic())).To(Equal(address.GetAddress()))
	})
})
