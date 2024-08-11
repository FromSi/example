package mappers_test

import (
	. "github.com/fromsi/example/internal/app/apiserver/domain/entities"
	. "github.com/fromsi/example/internal/app/apiserver/infrastructure/mappers"
	. "github.com/fromsi/example/internal/app/apiserver/infrastructure/models"
	"github.com/go-faker/faker/v4"
	"github.com/go-faker/faker/v4/pkg/options"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"gorm.io/gorm"
	"math/rand"
	"time"
)

var _ = Describe("Post", func() {
	var text string

	textLength := uint(rand.Intn(TextPostMaxLength-TextPostMinLength+1) + TextPostMinLength)

	Describe("Gorm", func() {
		It("can transform an nil into nil", func() {
			gormPostOne, err := ArrayEntityToArrayGorm(nil)

			Expect(err).NotTo(HaveOccurred())
			Expect(gormPostOne).To(BeNil())

			gormPostTwo, err := EntityToGorm(nil)

			Expect(err).NotTo(HaveOccurred())
			Expect(gormPostTwo).To(BeNil())

			entityPostOne, err := GormToEntity(nil)

			Expect(err).NotTo(HaveOccurred())
			Expect(entityPostOne).To(BeNil())

			entityPostTwo, err := ArrayGormToArrayEntity(nil)

			Expect(err).NotTo(HaveOccurred())
			Expect(entityPostTwo).To(BeNil())
		})

		It("can transform an entity into a model", func() {
			_ = faker.FakeData(&text, options.WithRandomStringLength(textLength))
			timeNow := time.Now()
			entityOne, err := NewPost(faker.UUIDHyphenated(), text, &timeNow, &timeNow, &timeNow)

			Expect(err).NotTo(HaveOccurred())

			_ = faker.FakeData(&text, options.WithRandomStringLength(textLength))
			entityTwo, err := NewPost(faker.UUIDHyphenated(), text, nil, nil, nil)

			Expect(err).NotTo(HaveOccurred())

			gormPost, err := EntityToGorm(entityOne)

			Expect(err).NotTo(HaveOccurred())
			Expect(gormPost).NotTo(BeNil())

			_, err = EntityToGorm(entityTwo)

			Expect(err).NotTo(HaveOccurred())

			Expect(gormPost.ID).To(Equal(entityOne.ID.GetId()))
			Expect(gormPost.Text).To(Equal(entityOne.Text.GetText()))
			Expect(gormPost.CreatedAt).To(Equal(*entityOne.CreatedAt))
			Expect(gormPost.UpdatedAt).To(Equal(*entityOne.UpdatedAt))
			Expect(gormPost.DeletedAt.Time).To(Equal(*entityOne.DeletedAt))
			Expect(gormPost.DeletedAt.Valid).To(BeTrue())
		})

		It("can transform an entity array into a model array", func() {
			arrayEntityPost := []Post{}
			arrayGormPost, err := ArrayEntityToArrayGorm(&arrayEntityPost)

			Expect(err).NotTo(HaveOccurred())
			Expect(arrayGormPost).NotTo(BeNil())
			Expect(*arrayGormPost).To(BeEmpty())

			_ = faker.FakeData(&text, options.WithRandomStringLength(textLength))
			entityOne, err := NewPost(faker.UUIDHyphenated(), text, nil, nil, nil)

			Expect(err).NotTo(HaveOccurred())

			_ = faker.FakeData(&text, options.WithRandomStringLength(textLength))
			timeNow := time.Now()
			entityTwo, err := NewPost(faker.UUIDHyphenated(), text, &timeNow, &timeNow, &timeNow)

			Expect(err).NotTo(HaveOccurred())

			arrayEntityPost = []Post{*entityOne, *entityTwo}
			arrayGormPost, err = ArrayEntityToArrayGorm(&arrayEntityPost)

			Expect(err).NotTo(HaveOccurred())
			Expect(arrayGormPost).NotTo(BeNil())

			arrayEntityPost = []Post{}
			_, err = ArrayEntityToArrayGorm(&arrayEntityPost)

			Expect(err).NotTo(HaveOccurred())

			gormPost := (*arrayGormPost)[0]

			Expect(gormPost.ID).To(Equal(entityOne.ID.GetId()))
			Expect(gormPost.Text).To(Equal(entityOne.Text.GetText()))
			Expect(gormPost.CreatedAt).NotTo(BeNil())
			Expect(gormPost.UpdatedAt).NotTo(BeNil())
			Expect(gormPost.DeletedAt).To(BeNil())

			gormPost = (*arrayGormPost)[1]

			Expect(gormPost.ID).To(Equal(entityTwo.ID.GetId()))
			Expect(gormPost.Text).To(Equal(entityTwo.Text.GetText()))
			Expect(gormPost.CreatedAt).To(Equal(*entityTwo.CreatedAt))
			Expect(gormPost.UpdatedAt).To(Equal(*entityTwo.UpdatedAt))
			Expect(gormPost.DeletedAt.Time).To(Equal(*entityTwo.DeletedAt))
			Expect(gormPost.DeletedAt.Valid).To(BeTrue())
		})

		It("can transform a model into an entity", func() {
			_ = faker.FakeData(&text, options.WithRandomStringLength(textLength))
			timeNow := time.Now()
			gormOne := GormPostModel{
				ID:        faker.UUIDHyphenated(),
				Text:      text,
				CreatedAt: timeNow,
				UpdatedAt: timeNow,
				DeletedAt: nil,
			}

			_ = faker.FakeData(&text, options.WithRandomStringLength(textLength))
			timeNow = time.Now()
			gormTwo := GormPostModel{
				ID:        faker.UUIDHyphenated(),
				Text:      text,
				CreatedAt: timeNow,
				UpdatedAt: timeNow,
				DeletedAt: &gorm.DeletedAt{
					Time:  timeNow,
					Valid: true,
				},
			}

			entityPostOne, err := GormToEntity(&gormOne)

			Expect(err).NotTo(HaveOccurred())
			Expect(entityPostOne).NotTo(BeNil())

			_, err = GormToEntity(&gormTwo)

			Expect(err).NotTo(HaveOccurred())

			Expect(entityPostOne.ID.GetId()).To(Equal(gormOne.ID))
			Expect(entityPostOne.Text.GetText()).To(Equal(gormOne.Text))
			Expect(entityPostOne.CreatedAt).To(Equal(&gormOne.CreatedAt))
			Expect(entityPostOne.UpdatedAt).To(Equal(&gormOne.UpdatedAt))
			Expect(entityPostOne.DeletedAt).To(BeNil())
		})

		It("can transform a model array into an entity array", func() {
			arrayGormPost := []GormPostModel{}
			arrayEntityPost, err := ArrayGormToArrayEntity(&arrayGormPost)

			Expect(err).NotTo(HaveOccurred())
			Expect(arrayEntityPost).NotTo(BeNil())
			Expect(*arrayEntityPost).To(BeEmpty())

			_ = faker.FakeData(&text, options.WithRandomStringLength(textLength))
			timeNow := time.Now()
			gormOne := GormPostModel{
				ID:        faker.UUIDHyphenated(),
				Text:      text,
				CreatedAt: timeNow,
				UpdatedAt: timeNow,
				DeletedAt: nil,
			}

			_ = faker.FakeData(&text, options.WithRandomStringLength(textLength))
			timeNow = time.Now()
			gormTwo := GormPostModel{
				ID:        faker.UUIDHyphenated(),
				Text:      text,
				CreatedAt: timeNow,
				UpdatedAt: timeNow,
				DeletedAt: &gorm.DeletedAt{
					Time:  timeNow,
					Valid: true,
				},
			}

			arrayGormPost = []GormPostModel{gormOne, gormTwo}
			arrayEntityPost, err = ArrayGormToArrayEntity(&arrayGormPost)

			Expect(err).NotTo(HaveOccurred())
			Expect(arrayEntityPost).NotTo(BeNil())

			arrayGormPost = []GormPostModel{}
			_, err = ArrayGormToArrayEntity(&arrayGormPost)

			Expect(err).NotTo(HaveOccurred())

			entityPost := (*arrayEntityPost)[0]

			Expect(entityPost.ID.GetId()).To(Equal(gormOne.ID))
			Expect(entityPost.Text.GetText()).To(Equal(gormOne.Text))
			Expect(entityPost.CreatedAt).To(Equal(&gormOne.CreatedAt))
			Expect(entityPost.UpdatedAt).To(Equal(&gormOne.UpdatedAt))
			Expect(entityPost.DeletedAt).To(BeNil())

			entityPost = (*arrayEntityPost)[1]

			Expect(entityPost.ID.GetId()).To(Equal(gormTwo.ID))
			Expect(entityPost.Text.GetText()).To(Equal(gormTwo.Text))
			Expect(entityPost.CreatedAt).To(Equal(&gormTwo.CreatedAt))
			Expect(entityPost.UpdatedAt).To(Equal(&gormTwo.UpdatedAt))
			Expect(entityPost.DeletedAt).To(Equal(&gormTwo.DeletedAt.Time))
		})
	})
})
