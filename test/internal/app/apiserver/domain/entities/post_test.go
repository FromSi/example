package entities_test

import (
	. "github.com/fromsi/example/internal/app/apiserver/domain/entities"
	"github.com/go-faker/faker/v4"
	"github.com/go-faker/faker/v4/pkg/options"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"math/rand"
	"time"
)

var _ = Describe("Post", func() {
	var err error

	It("can make a post", func() {
		id := faker.UUIDHyphenated()

		var text string
		textLength := uint(rand.Intn(TextMaxLength-TextMinLength+1) + TextMinLength)
		_ = faker.FakeData(&text, options.WithRandomStringLength(textLength))

		timeNow := time.Now()
		post, err := NewPost(id, text, &timeNow, &timeNow, &timeNow)

		Expect(err).NotTo(HaveOccurred())
		Expect(post).NotTo(BeNil())

		Expect(post.ID.GetId()).To(Equal(id))
		Expect(post.Text.GetText()).To(Equal(text))
		Expect(post.CreatedAt).To(Equal(&timeNow))
		Expect(post.UpdatedAt).To(Equal(&timeNow))
		Expect(post.DeletedAt).To(Equal(&timeNow))

		timeNow = time.Now()
		_, err = NewPost(id, text, &timeNow, &timeNow, &timeNow)

		Expect(err).NotTo(HaveOccurred())

		Expect(post.CreatedAt).NotTo(Equal(&timeNow))
		Expect(post.UpdatedAt).NotTo(Equal(&timeNow))
		Expect(post.DeletedAt).NotTo(Equal(&timeNow))
	})

	It("can change the text of the post", func() {
		id := faker.UUIDHyphenated()

		textLength := uint(rand.Intn(TextMaxLength-TextMinLength+1) + TextMinLength)

		var text string
		_ = faker.FakeData(&text, options.WithRandomStringLength(textLength))

		post, err := NewPost(id, text, nil, nil, nil)

		Expect(err).NotTo(HaveOccurred())
		Expect(post).NotTo(BeNil())

		var textTwo string
		_ = faker.FakeData(&textTwo, options.WithRandomStringLength(textLength))

		err = post.SetText(textTwo)

		Expect(post.Text.GetText()).To(Equal(textTwo))
	})

	Describe("ID Value Object", func() {
		var id *Id
		var value string

		It("can make a id", func() {
			id, err = NewId("")

			Expect(err).To(HaveOccurred())
			Expect(id).To(BeNil())

			id, err = NewId(faker.Word())

			Expect(err).To(HaveOccurred())
			Expect(id).To(BeNil())

			value = faker.UUIDHyphenated()
			id, err = NewId(value)

			Expect(err).NotTo(HaveOccurred())
			Expect(id).NotTo(BeNil())
			Expect(id.GetId()).To(Equal(value))
		})
	})

	Describe("Text Value Object", func() {
		var text *Text
		var value string

		It("can make a text", func() {
			text, err = NewText("")

			Expect(err).To(HaveOccurred())
			Expect(text).To(BeNil())

			_ = faker.FakeData(&value, options.WithRandomStringLength(TextMinLength-1))
			text, err = NewText(value)

			Expect(err).To(HaveOccurred())
			Expect(text).To(BeNil())

			_ = faker.FakeData(&value, options.WithRandomStringLength(TextMaxLength+1))
			text, err = NewText(value)

			Expect(err).To(HaveOccurred())
			Expect(text).To(BeNil())

			textLength := uint(rand.Intn(TextMaxLength-TextMinLength+1) + TextMinLength)
			_ = faker.FakeData(&value, options.WithRandomStringLength(textLength))
			text, err = NewText(value)

			Expect(err).NotTo(HaveOccurred())
			Expect(text).NotTo(BeNil())
			Expect(text.GetText()).To(Equal(value))
		})
	})
})
