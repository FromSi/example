package mappers_test

import (
	. "github.com/fromsi/example/internal/app/apiserver/application/cqrs/mappers"
	"github.com/fromsi/example/internal/pkg/tools"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Auth", func() {
	It("can transform an entity into ToGetMnemonicAuthQueryResponse", func() {
		addressBTC := tools.NewAddressBTC()

		response, err := ToGetMnemonicAuthQueryResponse(addressBTC.GetMnemonic())

		Expect(err).NotTo(HaveOccurred())

		Expect(response.Mnemonic).To(Equal(addressBTC.GetMnemonic()))
	})

	It("can transform an entity into ToGetAddressFromMnemonicAuthQueryResponse", func() {
		addressBTC := tools.NewAddressBTC()

		response, err := ToGetAddressFromMnemonicAuthQueryResponse(addressBTC.GetAddress())

		Expect(err).NotTo(HaveOccurred())

		Expect(response.Address).To(Equal(addressBTC.GetAddress()))
	})
})
