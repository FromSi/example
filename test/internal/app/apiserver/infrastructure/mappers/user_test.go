package mappers_test

import (
	. "github.com/fromsi/example/internal/app/apiserver/domain/entities"
	. "github.com/fromsi/example/internal/app/apiserver/infrastructure/mappers"
	. "github.com/fromsi/example/internal/app/apiserver/infrastructure/models"
	"github.com/fromsi/example/internal/pkg/tools"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"gorm.io/gorm"
	"time"
)

var _ = Describe("User", func() {
	Describe("Gorm", func() {
		It("can transform an nil into nil", func() {
			gormUserOne, err := ArrayEntityToArrayGormUser(nil)

			Expect(err).NotTo(HaveOccurred())
			Expect(gormUserOne).To(BeNil())

			gormUserTwo, err := EntityToGormUser(nil)

			Expect(err).NotTo(HaveOccurred())
			Expect(gormUserTwo).To(BeNil())

			entityUserOne, err := GormToEntityUser(nil)

			Expect(err).NotTo(HaveOccurred())
			Expect(entityUserOne).To(BeNil())

			entityUserTwo, err := ArrayGormToArrayEntityUser(nil)

			Expect(err).NotTo(HaveOccurred())
			Expect(entityUserTwo).To(BeNil())
		})

		It("can transform an entity into a model", func() {
			timeNow := time.Now()
			entityOne, err := NewUser(tools.NewAddressBTC().GetAddress(), &timeNow, &timeNow, &timeNow)

			Expect(err).NotTo(HaveOccurred())

			entityTwo, err := NewUser(tools.NewAddressBTC().GetAddress(), nil, nil, nil)

			Expect(err).NotTo(HaveOccurred())

			gormUser, err := EntityToGormUser(entityOne)

			Expect(err).NotTo(HaveOccurred())
			Expect(gormUser).NotTo(BeNil())

			_, err = EntityToGormUser(entityTwo)

			Expect(err).NotTo(HaveOccurred())

			Expect(gormUser.ID).To(Equal(entityOne.ID.GetId()))
			Expect(gormUser.CreatedAt).To(Equal(*entityOne.CreatedAt))
			Expect(gormUser.UpdatedAt).To(Equal(*entityOne.UpdatedAt))
			Expect(gormUser.DeletedAt.Time).To(Equal(*entityOne.DeletedAt))
			Expect(gormUser.DeletedAt.Valid).To(BeTrue())
		})

		It("can transform an entity array into a model array", func() {
			arrayEntityUser := []User{}
			arrayGormUser, err := ArrayEntityToArrayGormUser(&arrayEntityUser)

			Expect(err).NotTo(HaveOccurred())
			Expect(arrayGormUser).NotTo(BeNil())
			Expect(*arrayGormUser).To(BeEmpty())

			entityOne, err := NewUser(tools.NewAddressBTC().GetAddress(), nil, nil, nil)

			Expect(err).NotTo(HaveOccurred())

			timeNow := time.Now()
			entityTwo, err := NewUser(tools.NewAddressBTC().GetAddress(), &timeNow, &timeNow, &timeNow)

			Expect(err).NotTo(HaveOccurred())

			arrayEntityUser = []User{*entityOne, *entityTwo}
			arrayGormUser, err = ArrayEntityToArrayGormUser(&arrayEntityUser)

			Expect(err).NotTo(HaveOccurred())
			Expect(arrayGormUser).NotTo(BeNil())

			arrayEntityUser = []User{}
			_, err = ArrayEntityToArrayGormUser(&arrayEntityUser)

			Expect(err).NotTo(HaveOccurred())

			gormUser := (*arrayGormUser)[0]

			Expect(gormUser.ID).To(Equal(entityOne.ID.GetId()))
			Expect(gormUser.CreatedAt).NotTo(BeNil())
			Expect(gormUser.UpdatedAt).NotTo(BeNil())
			Expect(gormUser.DeletedAt).To(BeNil())

			gormUser = (*arrayGormUser)[1]

			Expect(gormUser.ID).To(Equal(entityTwo.ID.GetId()))
			Expect(gormUser.CreatedAt).To(Equal(*entityTwo.CreatedAt))
			Expect(gormUser.UpdatedAt).To(Equal(*entityTwo.UpdatedAt))
			Expect(gormUser.DeletedAt.Time).To(Equal(*entityTwo.DeletedAt))
			Expect(gormUser.DeletedAt.Valid).To(BeTrue())
		})

		It("can transform a model into an entity", func() {
			timeNow := time.Now()
			gormOne := GormUserModel{
				ID:        tools.NewAddressBTC().GetAddress(),
				CreatedAt: timeNow,
				UpdatedAt: timeNow,
				DeletedAt: nil,
			}

			timeNow = time.Now()
			gormTwo := GormUserModel{
				ID:        tools.NewAddressBTC().GetAddress(),
				CreatedAt: timeNow,
				UpdatedAt: timeNow,
				DeletedAt: &gorm.DeletedAt{
					Time:  timeNow,
					Valid: true,
				},
			}

			entityUserOne, err := GormToEntityUser(&gormOne)

			Expect(err).NotTo(HaveOccurred())
			Expect(entityUserOne).NotTo(BeNil())

			_, err = GormToEntityUser(&gormTwo)

			Expect(err).NotTo(HaveOccurred())

			Expect(entityUserOne.ID.GetId()).To(Equal(gormOne.ID))
			Expect(entityUserOne.CreatedAt).To(Equal(&gormOne.CreatedAt))
			Expect(entityUserOne.UpdatedAt).To(Equal(&gormOne.UpdatedAt))
			Expect(entityUserOne.DeletedAt).To(BeNil())
		})

		It("can transform a model array into an entity array", func() {
			arrayGormUser := []GormUserModel{}
			arrayEntityUser, err := ArrayGormToArrayEntityUser(&arrayGormUser)

			Expect(err).NotTo(HaveOccurred())
			Expect(arrayEntityUser).NotTo(BeNil())
			Expect(*arrayEntityUser).To(BeEmpty())

			timeNow := time.Now()
			gormOne := GormUserModel{
				ID:        tools.NewAddressBTC().GetAddress(),
				CreatedAt: timeNow,
				UpdatedAt: timeNow,
				DeletedAt: nil,
			}

			timeNow = time.Now()
			gormTwo := GormUserModel{
				ID:        tools.NewAddressBTC().GetAddress(),
				CreatedAt: timeNow,
				UpdatedAt: timeNow,
				DeletedAt: &gorm.DeletedAt{
					Time:  timeNow,
					Valid: true,
				},
			}

			arrayGormUser = []GormUserModel{gormOne, gormTwo}
			arrayEntityUser, err = ArrayGormToArrayEntityUser(&arrayGormUser)

			Expect(err).NotTo(HaveOccurred())
			Expect(arrayEntityUser).NotTo(BeNil())

			arrayGormUser = []GormUserModel{}
			_, err = ArrayGormToArrayEntityUser(&arrayGormUser)

			Expect(err).NotTo(HaveOccurred())

			entityUser := (*arrayEntityUser)[0]

			Expect(entityUser.ID.GetId()).To(Equal(gormOne.ID))
			Expect(entityUser.CreatedAt).To(Equal(&gormOne.CreatedAt))
			Expect(entityUser.UpdatedAt).To(Equal(&gormOne.UpdatedAt))
			Expect(entityUser.DeletedAt).To(BeNil())

			entityUser = (*arrayEntityUser)[1]

			Expect(entityUser.ID.GetId()).To(Equal(gormTwo.ID))
			Expect(entityUser.CreatedAt).To(Equal(&gormTwo.CreatedAt))
			Expect(entityUser.UpdatedAt).To(Equal(&gormTwo.UpdatedAt))
			Expect(entityUser.DeletedAt).To(Equal(&gormTwo.DeletedAt.Time))
		})
	})
})
