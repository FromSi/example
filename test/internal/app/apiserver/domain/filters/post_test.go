package filters_test

import (
	"github.com/fromsi/example/internal/app/apiserver/domain/filters"
	"github.com/go-faker/faker/v4"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Post", func() {
	It("Find Post Filter", func() {
		idOne := faker.UUIDHyphenated()
		idTwo := faker.UUIDHyphenated()

		filterOne, err := filters.NewFindPostFilter(idOne)

		Expect(err).NotTo(HaveOccurred())
		Expect(filterOne).NotTo(BeNil())

		filterTwo, err := filters.NewFindPostFilter(idTwo)

		Expect(err).NotTo(HaveOccurred())
		Expect(filterTwo).NotTo(BeNil())

		Expect(filterOne).NotTo(BeNil())
		Expect(filterOne.ID).To(Equal(idOne))

		Expect(filterTwo).NotTo(BeNil())
		Expect(filterTwo.ID).To(Equal(idTwo))
	})
})
