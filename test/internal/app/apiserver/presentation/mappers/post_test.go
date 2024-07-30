package mappers_test

import (
	. "github.com/fromsi/example/internal/app/apiserver/application/cqrs/responses"
	. "github.com/fromsi/example/internal/app/apiserver/domain/entities"
	. "github.com/fromsi/example/internal/app/apiserver/presentation/mappers"
	. "github.com/fromsi/example/internal/app/apiserver/presentation/responses"
	"github.com/go-faker/faker/v4"
	"github.com/go-faker/faker/v4/pkg/options"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"math/rand"
	"time"
)

var _ = Describe("Post", func() {
	var text string

	textLength := uint(rand.Intn(TextMaxLength-TextMinLength+1) + TextMinLength)

	Describe("Gin", func() {
		It("can transform an nil into nil", func() {
			queryResponseOne, err := ToGinShowPostResponse(nil)

			Expect(err).NotTo(HaveOccurred())
			Expect(queryResponseOne).To(BeNil())

			queryResponseTwo, err := ToGinIndexPostResponse(nil)

			Expect(err).NotTo(HaveOccurred())
			Expect(queryResponseTwo).To(BeNil())
		})

		It("can transform a query response into a response", func() {
			_ = faker.FakeData(&text, options.WithRandomStringLength(textLength))
			timeNow := time.Now()
			queryResponseOne := CqrsFindByIdQueryResponse{
				Data: QueryResponse{
					ID:        faker.UUIDHyphenated(),
					Text:      text,
					CreatedAt: &timeNow,
					UpdatedAt: &timeNow,
				},
				IsDeleted: false,
			}

			_ = faker.FakeData(&text, options.WithRandomStringLength(textLength))
			queryResponseTwo := CqrsFindByIdQueryResponse{
				Data: QueryResponse{
					ID:        faker.UUIDHyphenated(),
					Text:      text,
					CreatedAt: nil,
					UpdatedAt: nil,
				},
				IsDeleted: true,
			}

			response, err := ToGinShowPostResponse(&queryResponseOne)

			Expect(err).NotTo(HaveOccurred())
			Expect(response).NotTo(BeNil())

			_, err = ToGinShowPostResponse(&queryResponseTwo)

			Expect(err).NotTo(HaveOccurred())

			data, exists := response.Data.(PostResponse)

			Expect(exists).To(BeTrue())

			Expect(data.ID).To(Equal(queryResponseOne.Data.ID))
			Expect(data.Text).To(Equal(queryResponseOne.Data.Text))
			Expect(data.CreatedAt).To(Equal(queryResponseOne.Data.CreatedAt))
			Expect(data.UpdatedAt).To(Equal(queryResponseOne.Data.UpdatedAt))
		})

		It("can transform a query response into a response", func() {
			_ = faker.FakeData(&text, options.WithRandomStringLength(textLength))
			timeNow := time.Now()
			dataResponseOne := QueryResponse{
				ID:        faker.UUIDHyphenated(),
				Text:      text,
				CreatedAt: &timeNow,
				UpdatedAt: &timeNow,
			}

			_ = faker.FakeData(&text, options.WithRandomStringLength(textLength))
			dataResponseTwo := QueryResponse{
				ID:        faker.UUIDHyphenated(),
				Text:      text,
				CreatedAt: nil,
				UpdatedAt: nil,
			}

			pageableOne, _ := NewEntityPageable(MinPageOrder, MaxLimitItems, rand.Intn(100)+MinTotalItems)
			pageableTwo, _ := NewEntityPageable(rand.Intn(100)+MinPageOrder, MinLimitItems, MinTotalItems)

			responseOne := CqrsGetAllQueryResponse{
				Data:     []QueryResponse{dataResponseOne, dataResponseTwo},
				Pageable: pageableOne,
			}

			responseTwo := CqrsGetAllQueryResponse{
				Data:     []QueryResponse{},
				Pageable: pageableTwo,
			}

			response, err := ToGinIndexPostResponse(&responseOne)

			Expect(err).NotTo(HaveOccurred())
			Expect(response).NotTo(BeNil())

			_, err = ToGinIndexPostResponse(&responseTwo)

			Expect(err).NotTo(HaveOccurred())

			Expect(response.Pageable.Page).To(Equal(responseOne.Pageable.GetPage()))
			Expect(response.Pageable.Total).To(Equal(responseOne.Pageable.GetTotal()))
			Expect(response.Pageable.Limit).To(Equal(responseOne.Pageable.GetLimit()))

			responseData, exists := response.Data.([]PostResponse)

			Expect(exists).To(BeTrue())

			responseItem := responseData[0]

			Expect(responseItem.ID).To(Equal(dataResponseOne.ID))
			Expect(responseItem.Text).To(Equal(dataResponseOne.Text))
			Expect(responseItem.CreatedAt).To(Equal(dataResponseOne.CreatedAt))
			Expect(responseItem.UpdatedAt).To(Equal(dataResponseOne.UpdatedAt))

			responseItem = responseData[1]

			Expect(responseItem.ID).To(Equal(dataResponseTwo.ID))
			Expect(responseItem.Text).To(Equal(dataResponseTwo.Text))
			Expect(responseItem.CreatedAt).To(BeNil())
			Expect(responseItem.UpdatedAt).To(BeNil())
		})
	})
})
