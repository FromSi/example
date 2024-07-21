package entities_test

import (
	. "github.com/fromsi/example/internal/app/apiserver/domain/entities"
	"github.com/fromsi/example/internal/pkg/tools"
	"github.com/go-faker/faker/v4"
	"github.com/go-faker/faker/v4/pkg/options"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"reflect"
)

var _ = Describe("Sortable", func() {
	var err error

	Describe("Entity Sortable", func() {
		var entitySortable *EntitySortable
		var valueString string
		var value map[string]string

		It("must support the sortable interface", func() {
			instance := EntitySortable{}

			myInterfaceType := reflect.TypeOf((*Sortable)(nil)).Elem()

			for i := 0; i < myInterfaceType.NumMethod(); i++ {
				method := myInterfaceType.Method(i)
				_, ok := reflect.TypeOf(&instance).MethodByName(method.Name)

				Expect(ok).To(BeTrue(), "Method '%s' not implemented", method.Name)
			}
		})

		It("cannot take a value different from desc or asc", func() {
			_ = faker.FakeData(&valueString, options.WithRandomStringLength(4))

			if valueString == OrderAsc || valueString == OrderDesc {
				valueString = "order"
			}

			value = map[string]string{"field": valueString}
			entitySortable, err = NewEntitySortable(value)

			Expect(err).To(HaveOccurred())
		})

		It("can take the value asc", func() {
			value = map[string]string{"field": OrderAsc}
			entitySortable, err = NewEntitySortable(value)

			Expect(err).NotTo(HaveOccurred())
			Expect(entitySortable.Data).To(Equal(value))
		})

		It("can take the value desc", func() {
			value = map[string]string{"field": OrderDesc}
			entitySortable, err = NewEntitySortable(value)

			Expect(err).NotTo(HaveOccurred())
			Expect(entitySortable.Data).To(Equal(value))
		})

		It("can extract the iterator", func() {
			value = map[string]string{"field": OrderDesc}
			entitySortable, err = NewEntitySortable(value)

			Expect(err).NotTo(HaveOccurred())
			Expect(entitySortable.GetIterator()).To(Equal(tools.NewMapStringIterator(value)))
		})
	})
})
