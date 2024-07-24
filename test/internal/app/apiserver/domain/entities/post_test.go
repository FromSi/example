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
	var textOne string
	var textTwo string
	var err error

	textLength := uint(rand.Intn(TextMaxLength-TextMinLength+1) + TextMinLength)

	It("can make a post", func() {
		idOne := faker.UUIDHyphenated()
		_ = faker.FakeData(&textOne, options.WithRandomStringLength(textLength))
		timeNowOne := time.Now()
		post, err := NewPost(idOne, textOne, &timeNowOne, &timeNowOne, &timeNowOne)

		Expect(err).NotTo(HaveOccurred())
		Expect(post).NotTo(BeNil())

		idTwo := faker.UUIDHyphenated()
		_ = faker.FakeData(&textTwo, options.WithRandomStringLength(textLength))
		timeNowTwo := time.Now()
		_, err = NewPost(idTwo, textTwo, &timeNowTwo, &timeNowTwo, &timeNowTwo)

		Expect(err).NotTo(HaveOccurred())

		Expect(post.ID.GetId()).To(Equal(idOne))
		Expect(post.Text.GetText()).To(Equal(textOne))
		Expect(post.CreatedAt).To(Equal(&timeNowOne))
		Expect(post.UpdatedAt).To(Equal(&timeNowOne))
		Expect(post.DeletedAt).To(Equal(&timeNowOne))
	})

	It("can change the textOne of the post", func() {
		id := faker.UUIDHyphenated()
		_ = faker.FakeData(&textOne, options.WithRandomStringLength(textLength))
		post, err := NewPost(id, textOne, nil, nil, nil)

		Expect(err).NotTo(HaveOccurred())
		Expect(post).NotTo(BeNil())

		_ = faker.FakeData(&textTwo, options.WithRandomStringLength(textLength))

		err = post.SetText(textTwo)

		Expect(err).NotTo(HaveOccurred())

		Expect(post.Text.GetText()).To(Equal(textTwo))
	})

	Describe("ID Value Object", func() {
		var idValueObject *Id

		It("can make a id", func() {
			idValueObject, err = NewId("")

			Expect(err).To(HaveOccurred())
			Expect(idValueObject).To(BeNil())

			idValueObject, err = NewId(faker.Word())

			Expect(err).To(HaveOccurred())
			Expect(idValueObject).To(BeNil())

			textOne = faker.UUIDHyphenated()
			idValueObject, err = NewId(textOne)

			Expect(err).NotTo(HaveOccurred())
			Expect(idValueObject).NotTo(BeNil())

			textTwo = faker.UUIDHyphenated()
			_, err = NewId(textTwo)

			Expect(err).NotTo(HaveOccurred())

			Expect(idValueObject.GetId()).To(Equal(textOne))
		})
	})

	Describe("Text Value Object", func() {
		var textValueObject *Text

		It("can make a textOne", func() {
			textValueObject, err = NewText("")

			Expect(err).To(HaveOccurred())
			Expect(textValueObject).To(BeNil())

			_ = faker.FakeData(&textOne, options.WithRandomStringLength(TextMinLength-1))
			textValueObject, err = NewText(textOne)

			Expect(err).To(HaveOccurred())
			Expect(textValueObject).To(BeNil())

			_ = faker.FakeData(&textOne, options.WithRandomStringLength(TextMaxLength+1))
			textValueObject, err = NewText(textOne)

			Expect(err).To(HaveOccurred())
			Expect(textValueObject).To(BeNil())

			_ = faker.FakeData(&textOne, options.WithRandomStringLength(textLength))
			textValueObject, err = NewText(textOne)

			Expect(err).NotTo(HaveOccurred())
			Expect(textValueObject).NotTo(BeNil())

			_ = faker.FakeData(&textTwo, options.WithRandomStringLength(textLength))
			_, err = NewText(textTwo)

			Expect(err).NotTo(HaveOccurred())

			Expect(textValueObject.GetText()).To(Equal(textOne))
		})
	})
})
