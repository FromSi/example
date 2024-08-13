package filters_test

import (
	"github.com/fromsi/example/internal/app/apiserver/domain/filters"
	"github.com/fromsi/example/internal/pkg/tools"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("User", func() {
	It("Find User Filter", func() {
		addressOne := tools.NewAddressBTC()
		addressTwo := tools.NewAddressBTC()

		filterOne, err := filters.NewFindUserFilter(addressOne.GetAddress())

		Expect(err).NotTo(HaveOccurred())
		Expect(filterOne).NotTo(BeNil())

		filterTwo, err := filters.NewFindUserFilter(addressTwo.GetAddress())

		Expect(err).NotTo(HaveOccurred())
		Expect(filterTwo).NotTo(BeNil())

		Expect(filterOne).NotTo(BeNil())
		Expect(filterOne.ID).To(Equal(addressOne.GetAddress()))

		Expect(filterTwo).NotTo(BeNil())
		Expect(filterTwo.ID).To(Equal(addressTwo.GetAddress()))
	})
})
