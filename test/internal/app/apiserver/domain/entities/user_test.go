package entities_test

import (
	. "github.com/fromsi/example/internal/app/apiserver/domain/entities"
	"github.com/fromsi/example/internal/pkg/tools"
	"github.com/go-faker/faker/v4"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"time"
)

var _ = Describe("User", func() {
	var err error

	It("can make a user", func() {
		timeNowOne := time.Now()
		user, err := NewUser("", &timeNowOne, &timeNowOne, &timeNowOne)

		Expect(err).NotTo(HaveOccurred())
		Expect(user).NotTo(BeNil())

		idTwo := faker.UUIDHyphenated()
		timeNowTwo := time.Now()
		_, err = NewUser(idTwo, &timeNowTwo, &timeNowTwo, &timeNowTwo)

		Expect(err).NotTo(HaveOccurred())

		Expect(user.CreatedAt).To(Equal(&timeNowOne))
		Expect(user.UpdatedAt).To(Equal(&timeNowOne))
		Expect(user.DeletedAt).To(Equal(&timeNowOne))
	})

	Describe("ID Value Object", func() {
		var idValueObject *IdUser

		It("can make correctly an id", func() {
			idValueObject, err = NewIdUser("")

			Expect(err).NotTo(HaveOccurred())
			Expect(idValueObject).NotTo(BeNil())

			addressBTCOne := tools.NewAddressBTC()
			idValueObject, err = NewIdUser(addressBTCOne.GetAddress())

			Expect(err).NotTo(HaveOccurred())
			Expect(idValueObject).NotTo(BeNil())

			addressBTCTwo := tools.NewAddressBTC()
			_, err = NewIdUser(addressBTCTwo.GetAddress())

			Expect(err).NotTo(HaveOccurred())

			Expect(idValueObject.GetId()).To(Equal(addressBTCOne.GetAddress()))
		})
	})
})
