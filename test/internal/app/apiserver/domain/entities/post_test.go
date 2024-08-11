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
	textLength := uint(rand.Intn(TextPostMaxLength-TextPostMinLength+1) + TextPostMinLength)

	It("can make a post", func() {
		var textOne string
		var textTwo string

		_ = faker.FakeData(&textOne, options.WithRandomStringLength(textLength))
		timeNowOne := time.Now()
		post, err := NewPost("", textOne, &timeNowOne, &timeNowOne, &timeNowOne)

		Expect(err).NotTo(HaveOccurred())
		Expect(post).NotTo(BeNil())

		_ = faker.FakeData(&textTwo, options.WithRandomStringLength(textLength))
		timeNowTwo := time.Now()
		_, err = NewPost("", textTwo, &timeNowTwo, &timeNowTwo, &timeNowTwo)

		Expect(err).NotTo(HaveOccurred())

		Expect(post.Text.GetText()).To(Equal(textOne))
		Expect(post.CreatedAt).To(Equal(&timeNowOne))
		Expect(post.UpdatedAt).To(Equal(&timeNowOne))
		Expect(post.DeletedAt).To(Equal(&timeNowOne))
	})

	It("can change the text of the post", func() {
		var textOne string
		var textTwo string

		_ = faker.FakeData(&textOne, options.WithRandomStringLength(textLength))
		post, err := NewPost("", textOne, nil, nil, nil)

		Expect(err).NotTo(HaveOccurred())
		Expect(post).NotTo(BeNil())

		_ = faker.FakeData(&textTwo, options.WithRandomStringLength(textLength))

		err = post.SetText(textTwo)

		Expect(err).NotTo(HaveOccurred())

		Expect(post.Text.GetText()).To(Equal(textTwo))
	})

	Describe("ID Value Object", func() {
		var idValueObject *IdPost
		var err error

		It("can make correctly an id", func() {
			idValueObject, err = NewIdPost("")

			Expect(err).NotTo(HaveOccurred())
			Expect(idValueObject).NotTo(BeNil())

			idValueObject, err = NewIdPost(faker.Word())

			Expect(err).To(HaveOccurred())
			Expect(idValueObject).To(BeNil())

			idOne := faker.UUIDHyphenated()
			idValueObject, err = NewIdPost(idOne)

			Expect(err).NotTo(HaveOccurred())
			Expect(idValueObject).NotTo(BeNil())

			idTwo := faker.UUIDHyphenated()
			_, err = NewIdPost(idTwo)

			Expect(err).NotTo(HaveOccurred())

			Expect(idValueObject.GetId()).To(Equal(idOne))
		})
	})

	Describe("Text Value Object", func() {
		var textOne string
		var textTwo string
		var textValueObject *TextPost
		var err error

		It("can make correctly a text", func() {
			textValueObject, err = NewTextPost("")

			Expect(err).To(HaveOccurred())
			Expect(textValueObject).To(BeNil())

			_ = faker.FakeData(&textOne, options.WithRandomStringLength(TextPostMinLength-1))
			textValueObject, err = NewTextPost(textOne)

			Expect(err).To(HaveOccurred())
			Expect(textValueObject).To(BeNil())

			_ = faker.FakeData(&textOne, options.WithRandomStringLength(TextPostMaxLength+1))
			textValueObject, err = NewTextPost(textOne)

			Expect(err).To(HaveOccurred())
			Expect(textValueObject).To(BeNil())

			_ = faker.FakeData(&textOne, options.WithRandomStringLength(textLength))
			textValueObject, err = NewTextPost(textOne)

			Expect(err).NotTo(HaveOccurred())
			Expect(textValueObject).NotTo(BeNil())

			_ = faker.FakeData(&textTwo, options.WithRandomStringLength(textLength))
			_, err = NewTextPost(textTwo)

			Expect(err).NotTo(HaveOccurred())

			Expect(textValueObject.GetText()).To(Equal(textOne))
		})
	})
})
