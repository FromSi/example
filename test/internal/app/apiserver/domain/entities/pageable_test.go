package entities_test

import (
	. "github.com/fromsi/example/internal/app/apiserver/domain/entities"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"math/rand"
	"reflect"
)

const (
	LengthTotalItems = 100
)

var _ = Describe("Pageable", func() {
	var err error

	Describe("Entity Pageable", func() {
		var entityPageable *EntityPageable

		BeforeEach(func() {
			entityPageable, err = NewEntityPageable(MinPageOrder, MaxLimitItems, LengthTotalItems)

			Expect(err).NotTo(HaveOccurred())
			Expect(entityPageable).NotTo(BeNil())

			_, err = NewEntityPageable(rand.Intn(100)+MinPageOrder, MinLimitItems, MinTotalItems)

			Expect(err).NotTo(HaveOccurred())
		})

		It("must support the pageable interface", func() {
			instance := EntityPageable{}

			myInterfaceType := reflect.TypeOf((*Pageable)(nil)).Elem()

			for i := 0; i < myInterfaceType.NumMethod(); i++ {
				method := myInterfaceType.Method(i)
				_, ok := reflect.TypeOf(&instance).MethodByName(method.Name)

				Expect(ok).To(BeTrue(), "Method '%s' not implemented", method.Name)
			}
		})

		It("should change the page correctly", func() {
			Expect(entityPageable.GetPage()).To(Equal(MinPageOrder))

			err = entityPageable.SetPage(MinPageOrder - 1)

			Expect(err).NotTo(HaveOccurred())
			Expect(entityPageable.GetPage()).To(Equal(MinPageOrder))
		})

		It("should change the limit correctly", func() {
			Expect(entityPageable.GetLimit()).To(Equal(MaxLimitItems))

			err = entityPageable.SetLimit(MaxLimitItems + 1)

			Expect(err).NotTo(HaveOccurred())
			Expect(entityPageable.GetLimit()).To(Equal(MaxLimitItems))

			err = entityPageable.SetLimit(MinLimitItems - 1)

			Expect(err).NotTo(HaveOccurred())
			Expect(entityPageable.GetLimit()).To(Equal(MaxLimitItems))
		})

		It("should change the total correctly", func() {
			Expect(entityPageable.GetTotal()).To(Equal(LengthTotalItems))

			err = entityPageable.SetTotal(MinTotalItems - 1)

			Expect(err).NotTo(HaveOccurred())
			Expect(entityPageable.GetTotal()).To(Equal(MinTotalItems))
		})

		It("can correctly take the value of total pages", func() {
			totalPages := (LengthTotalItems + MaxLimitItems - 1) / MaxLimitItems
			Expect(entityPageable.GetTotalPages()).To(Equal(totalPages))

			err = entityPageable.SetTotal(0)

			Expect(err).NotTo(HaveOccurred())
			Expect(entityPageable.GetTotalPages()).To(Equal(MinPageOrder))
		})

		It("can correctly take the value of next page", func() {
			_ = entityPageable.SetPage(1)
			_ = entityPageable.SetLimit(10)

			Expect(entityPageable.GetNext()).To(Equal(2))

			_ = entityPageable.SetPage(-1)

			Expect(entityPageable.GetNext()).To(Equal(2))

			_ = entityPageable.SetPage(9)

			Expect(entityPageable.GetNext()).To(Equal(10))

			_ = entityPageable.SetPage(10)

			Expect(entityPageable.GetNext()).To(Equal(10))

			_ = entityPageable.SetPage(11)

			Expect(entityPageable.GetNext()).To(Equal(10))
		})

		It("can correctly take the value of prev page", func() {
			_ = entityPageable.SetPage(1)
			_ = entityPageable.SetLimit(10)

			Expect(entityPageable.GetPrev()).To(Equal(1))

			_ = entityPageable.SetPage(-1)

			Expect(entityPageable.GetPrev()).To(Equal(1))

			_ = entityPageable.SetPage(10)

			Expect(entityPageable.GetPrev()).To(Equal(9))
		})
	})
})
