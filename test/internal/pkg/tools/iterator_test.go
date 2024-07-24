package tools_test

import (
	. "github.com/fromsi/example/internal/pkg/tools"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Iterator", func() {
	Describe("Map String Iterator", func() {
		var key string
		var value string
		data := map[string]string{}

		var iterator *MapStringIterator

		It("can correctly ignore if empty", func() {
			iterator = NewMapStringIterator(data)

			Expect(iterator).NotTo(BeNil())
			Expect(iterator.HasNext()).To(BeFalse())

			key, value = iterator.GetNext()

			Expect(key).To(BeEmpty())
			Expect(value).To(BeEmpty())
		})

		It("can take data correctly if there is only one item", func() {
			data := map[string]string{
				"one": "valueOne",
			}

			iterator = NewMapStringIterator(data)

			Expect(iterator).NotTo(BeNil())
			Expect(iterator.HasNext()).To(BeTrue())

			key, value = iterator.GetNext()

			Expect(key).To(Equal("one"))
			Expect(value).To(Equal("valueOne"))

			Expect(iterator.HasNext()).To(BeFalse())

			key, value = iterator.GetNext()

			Expect(key).To(BeEmpty())
			Expect(value).To(BeEmpty())
		})

		It("can take data correctly if there is only three items", func() {
			data := map[string]string{
				"one":   "valueOne",
				"two":   "valueTwo",
				"three": "valueThree",
			}

			iterator = NewMapStringIterator(data)

			Expect(iterator).NotTo(BeNil())

			for i := 0; i < 3; i++ {
				Expect(iterator.HasNext()).To(BeTrue())

				key, value = iterator.GetNext()

				keyMatcher := Or(Equal("one"), Equal("two"), Equal("three"))
				valueMatcher := Or(Equal("valueOne"), Equal("valueTwo"), Equal("valueThree"))

				Expect(key).To(keyMatcher)
				Expect(value).To(valueMatcher)
			}

			Expect(iterator.HasNext()).To(BeFalse())

			key, value = iterator.GetNext()

			Expect(key).To(BeEmpty())
			Expect(value).To(BeEmpty())
		})
	})
})
