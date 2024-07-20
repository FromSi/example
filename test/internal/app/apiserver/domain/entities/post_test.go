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

		createdAt := time.Now()
		updatedAt := time.Now()
		deletedAt := time.Now()

		post, err := NewPost(id, text, &createdAt, &updatedAt, &deletedAt)

		Expect(err).NotTo(HaveOccurred())

		Expect(post.ID.GetId()).To(Equal(id))
		Expect(post.Text.GetText()).To(Equal(text))
		Expect(post.CreatedAt).To(Equal(&createdAt))
		Expect(post.UpdatedAt).To(Equal(&updatedAt))
		Expect(post.DeletedAt).To(Equal(&deletedAt))

		createdAt = time.Now()
		updatedAt = time.Now()
		deletedAt = time.Now()

		_, err = NewPost(id, text, &createdAt, &updatedAt, &deletedAt)

		Expect(err).NotTo(HaveOccurred())

		Expect(post.CreatedAt).NotTo(Equal(&createdAt))
		Expect(post.UpdatedAt).NotTo(Equal(&updatedAt))
		Expect(post.DeletedAt).NotTo(Equal(&deletedAt))
	})

	It("can change the text of the post", func() {
		id := faker.UUIDHyphenated()

		textLength := uint(rand.Intn(TextMaxLength-TextMinLength+1) + TextMinLength)

		var text string
		_ = faker.FakeData(&text, options.WithRandomStringLength(textLength))

		post, err := NewPost(id, text, nil, nil, nil)

		Expect(err).NotTo(HaveOccurred())

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

			id, err = NewId(faker.Word())

			Expect(err).To(HaveOccurred())

			value = faker.UUIDHyphenated()
			id, err = NewId(value)

			Expect(err).NotTo(HaveOccurred())
			Expect(id.GetId()).To(Equal(value))
		})
	})

	Describe("Text Value Object", func() {
		var text *Text
		var value string

		It("can make a text", func() {
			text, err = NewText("")

			Expect(err).To(HaveOccurred())

			_ = faker.FakeData(&value, options.WithRandomStringLength(TextMinLength-1))
			text, err = NewText(value)

			Expect(err).To(HaveOccurred())

			_ = faker.FakeData(&value, options.WithRandomStringLength(TextMaxLength+1))
			text, err = NewText(value)

			Expect(err).To(HaveOccurred())

			textLength := uint(rand.Intn(TextMaxLength-TextMinLength+1) + TextMinLength)
			_ = faker.FakeData(&value, options.WithRandomStringLength(textLength))
			text, err = NewText(value)

			Expect(err).NotTo(HaveOccurred())
			Expect(text.GetText()).To(Equal(value))
		})
	})
})
