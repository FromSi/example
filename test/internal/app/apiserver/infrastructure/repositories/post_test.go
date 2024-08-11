package repositories_test

import (
	"github.com/fromsi/example/cmd/apiserver/config"
	"github.com/fromsi/example/cmd/apiserver/database"
	. "github.com/fromsi/example/internal/app/apiserver/domain/entities"
	"github.com/fromsi/example/internal/app/apiserver/domain/filters"
	. "github.com/fromsi/example/internal/app/apiserver/infrastructure/repositories"
	"github.com/go-faker/faker/v4"
	"github.com/go-faker/faker/v4/pkg/options"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"math/rand"
	"reflect"
	"time"
)

var _ = Describe("Post", func() {
	var applicationConfig *config.Config
	var relationDatabase *database.RelationDatabase

	var err error

	var textOne string
	var textTwo string

	textLength := uint(rand.Intn(TextPostMaxLength-TextPostMinLength+1) + TextPostMinLength)

	BeforeEach(func() {
		applicationConfig, err = config.NewConfig()

		Expect(err).NotTo(HaveOccurred())

		relationDatabase, err = database.NewRelationDatabase(applicationConfig)

		Expect(err).NotTo(HaveOccurred())
	})

	Describe("Gorm", func() {
		var gormDatabase database.TestGormDB
		var postRepository *GormPostRepository

		BeforeEach(func() {
			gormDatabase, err = relationDatabase.GetTestGormORM()

			Expect(err).NotTo(HaveOccurred())

			postRepository = NewGormPostRepository(gormDatabase)
		})

		AfterEach(func() {
			err = postRepository.Truncate()

			Expect(err).NotTo(HaveOccurred())
		})

		It("must support the query repository interface", func() {
			instance := GormPostRepository{}

			myInterfaceType := reflect.TypeOf((*QueryRepository)(nil)).Elem()

			for i := 0; i < myInterfaceType.NumMethod(); i++ {
				method := myInterfaceType.Method(i)
				_, ok := reflect.TypeOf(&instance).MethodByName(method.Name)

				Expect(ok).To(BeTrue(), "Method '%s' not implemented", method.Name)
			}
		})

		It("must support the mutable repository interface", func() {
			instance := GormPostRepository{}

			myInterfaceType := reflect.TypeOf((*MutableRepository)(nil)).Elem()

			for i := 0; i < myInterfaceType.NumMethod(); i++ {
				method := myInterfaceType.Method(i)
				_, ok := reflect.TypeOf(&instance).MethodByName(method.Name)

				Expect(ok).To(BeTrue(), "Method '%s' not implemented", method.Name)
			}
		})

		It("can make a crud operations", func() {
			By("Creating")

			idOne := faker.UUIDHyphenated()
			_ = faker.FakeData(&textOne, options.WithRandomStringLength(textLength))
			timeNowOne := time.Now()
			postOne, err := NewPost(idOne, textOne, &timeNowOne, &timeNowOne, nil)

			Expect(err).NotTo(HaveOccurred())
			Expect(postOne).NotTo(BeNil())

			postFilterOne, err := filters.NewFindPostFilter(idOne)

			Expect(err).NotTo(HaveOccurred())

			postOneFromDatabase, err := postRepository.FindByFilterWithTrashed(*postFilterOne)

			Expect(err).To(HaveOccurred())
			Expect(postOneFromDatabase).To(BeNil())

			err = postRepository.Create(postOne)

			Expect(err).NotTo(HaveOccurred())

			By("Finding By Filter With Trashed")

			postOneFromDatabase, err = postRepository.FindByFilterWithTrashed(*postFilterOne)

			Expect(err).NotTo(HaveOccurred())
			Expect(postOneFromDatabase.ID.GetId()).To(Equal(postOne.ID.GetId()))

			By("Updating")

			_ = faker.FakeData(&textOne, options.WithRandomStringLength(textLength))

			err = postOne.SetText(textOne)

			Expect(err).NotTo(HaveOccurred())

			err = postRepository.UpdateById(idOne, postOne)

			Expect(err).NotTo(HaveOccurred())

			postOneFromDatabase, err = postRepository.FindByFilterWithTrashed(*postFilterOne)

			Expect(err).NotTo(HaveOccurred())
			Expect(postOneFromDatabase.ID.GetId()).To(Equal(postOne.ID.GetId()))
			Expect(postOneFromDatabase.Text.GetText()).To(Equal(postOne.Text.GetText()))

			By("Deleting")

			err = postRepository.DeleteById(idOne)

			Expect(err).NotTo(HaveOccurred())

			postOneFromDatabase, err = postRepository.FindByFilterWithTrashed(*postFilterOne)

			Expect(err).NotTo(HaveOccurred())
			Expect(postOneFromDatabase.ID.GetId()).To(Equal(postOne.ID.GetId()))
			Expect(postOneFromDatabase.DeletedAt).NotTo(BeNil())

			By("Restoring")

			err = postRepository.RestoreById(idOne)

			Expect(err).NotTo(HaveOccurred())

			postOneFromDatabase, err = postRepository.FindByFilterWithTrashed(*postFilterOne)

			Expect(err).NotTo(HaveOccurred())
			Expect(postOneFromDatabase.ID.GetId()).To(Equal(postOne.ID.GetId()))
			Expect(postOneFromDatabase.DeletedAt).To(BeNil())

			By("Getting Total")

			idTwo := faker.UUIDHyphenated()
			_ = faker.FakeData(&textTwo, options.WithRandomStringLength(textLength))
			timeNowTwo := time.Now()
			postTwo, err := NewPost(idTwo, textTwo, &timeNowTwo, &timeNowTwo, nil)

			Expect(err).NotTo(HaveOccurred())
			Expect(postTwo).NotTo(BeNil())

			err = postRepository.Create(postTwo)

			Expect(err).NotTo(HaveOccurred())

			postTotal, err := postRepository.GetTotal()

			Expect(err).NotTo(HaveOccurred())
			Expect(postTotal).To(Equal(2))

			By("Getting All")

			pageable, err := NewEntityPageable(MinPageOrder, MaxLimitItems, MinTotalItems)

			Expect(err).NotTo(HaveOccurred())

			sortable, err := NewEntitySortable(map[string]string{})

			Expect(err).NotTo(HaveOccurred())

			posts, err := postRepository.GetAll(pageable, sortable)

			Expect(err).NotTo(HaveOccurred())

			for _, post := range *posts {
				Expect(post.ID.GetId()).To(Or(Equal(idOne), Equal(idTwo)))
			}
		})
	})
})
