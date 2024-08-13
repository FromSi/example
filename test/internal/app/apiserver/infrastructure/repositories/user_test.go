package repositories_test

import (
	"github.com/fromsi/example/cmd/apiserver/config"
	"github.com/fromsi/example/cmd/apiserver/database"
	. "github.com/fromsi/example/internal/app/apiserver/domain/entities"
	"github.com/fromsi/example/internal/app/apiserver/domain/filters"
	. "github.com/fromsi/example/internal/app/apiserver/infrastructure/repositories"
	"github.com/fromsi/example/internal/pkg/tools"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"reflect"
)

var _ = Describe("User", func() {
	var applicationConfig *config.Config
	var relationDatabase *database.RelationDatabase

	var err error

	BeforeEach(func() {
		applicationConfig, err = config.NewConfig()

		Expect(err).NotTo(HaveOccurred())

		relationDatabase, err = database.NewRelationDatabase(applicationConfig)

		Expect(err).NotTo(HaveOccurred())
	})

	Describe("Gorm", func() {
		var gormDatabase database.TestGormDB
		var userRepository *GormUserRepository

		BeforeEach(func() {
			gormDatabase, err = relationDatabase.GetTestGormORM()

			Expect(err).NotTo(HaveOccurred())

			userRepository = NewGormUserRepository(gormDatabase)
		})

		AfterEach(func() {
			err = userRepository.Truncate()

			Expect(err).NotTo(HaveOccurred())
		})

		It("must support the query repository interface", func() {
			instance := GormUserRepository{}

			myInterfaceType := reflect.TypeOf((*QueryRepository)(nil)).Elem()

			for i := 0; i < myInterfaceType.NumMethod(); i++ {
				method := myInterfaceType.Method(i)
				_, ok := reflect.TypeOf(&instance).MethodByName(method.Name)

				Expect(ok).To(BeTrue(), "Method '%s' not implemented", method.Name)
			}
		})

		It("must support the mutable repository interface", func() {
			instance := GormUserRepository{}

			myInterfaceType := reflect.TypeOf((*MutableRepository)(nil)).Elem()

			for i := 0; i < myInterfaceType.NumMethod(); i++ {
				method := myInterfaceType.Method(i)
				_, ok := reflect.TypeOf(&instance).MethodByName(method.Name)

				Expect(ok).To(BeTrue(), "Method '%s' not implemented", method.Name)
			}
		})

		It("can make a crud operations", func() {
			By("Creating")

			idOne := tools.NewAddressBTC().GetAddress()
			err = userRepository.CreateIfNotExistsById(idOne)

			Expect(err).NotTo(HaveOccurred())

			By("Finding By Filter With Trashed")

			userFilterOne, err := filters.NewFindUserFilter(idOne)

			Expect(err).NotTo(HaveOccurred())

			userOneFromDatabase, err := userRepository.FindByFilterWithTrashed(*userFilterOne)

			Expect(err).NotTo(HaveOccurred())
			Expect(userOneFromDatabase.ID.GetId()).To(Equal(idOne))

			By("Deleting")

			err = userRepository.DeleteById(idOne)

			Expect(err).NotTo(HaveOccurred())

			userOneFromDatabase, err = userRepository.FindByFilterWithTrashed(*userFilterOne)

			Expect(err).NotTo(HaveOccurred())
			Expect(userOneFromDatabase.ID.GetId()).To(Equal(idOne))
			Expect(userOneFromDatabase.DeletedAt).NotTo(BeNil())

			By("Restoring")

			err = userRepository.RestoreById(idOne)

			Expect(err).NotTo(HaveOccurred())

			userOneFromDatabase, err = userRepository.FindByFilterWithTrashed(*userFilterOne)

			Expect(err).NotTo(HaveOccurred())
			Expect(userOneFromDatabase.ID.GetId()).To(Equal(idOne))
			Expect(userOneFromDatabase.DeletedAt).To(BeNil())

			By("Getting Total")

			idTwo := tools.NewAddressBTC().GetAddress()
			err = userRepository.CreateIfNotExistsById(idTwo)

			Expect(err).NotTo(HaveOccurred())

			userTotal, err := userRepository.GetTotal()

			Expect(err).NotTo(HaveOccurred())
			Expect(userTotal).To(Equal(2))

			By("Creating If Not Exists By Id")

			err = userRepository.CreateIfNotExistsById(idTwo)

			Expect(err).NotTo(HaveOccurred())

			userTotal, err = userRepository.GetTotal()

			Expect(err).NotTo(HaveOccurred())
			Expect(userTotal).To(Equal(2))

			By("Getting All")

			pageable, err := NewEntityPageable(MinPageOrder, MaxLimitItems, MinTotalItems)

			Expect(err).NotTo(HaveOccurred())

			sortable, err := NewEntitySortable(map[string]string{})

			Expect(err).NotTo(HaveOccurred())

			users, err := userRepository.GetAll(pageable, sortable)

			Expect(err).NotTo(HaveOccurred())

			for _, user := range *users {
				Expect(user.ID.GetId()).To(Or(Equal(idOne), Equal(idTwo)))
			}
		})
	})
})
