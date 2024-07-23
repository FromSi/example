package filters_test

import (
	"github.com/fromsi/example/internal/app/apiserver/domain/filters"
	"github.com/go-faker/faker/v4"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Post", func() {
	It("Find Post Filter", func() {
		id := faker.UUIDHyphenated()
		filter, err := filters.NewFindPostFilter(id)

		Expect(err).NotTo(HaveOccurred())
		Expect(filter.ID).To(Equal(id))
	})
})
