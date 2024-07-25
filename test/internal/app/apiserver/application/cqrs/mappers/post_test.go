package filters_test

import (
	. "github.com/fromsi/example/internal/app/apiserver/application/cqrs/mappers"
	. "github.com/fromsi/example/internal/app/apiserver/domain/entities"
	"github.com/go-faker/faker/v4"
	"github.com/go-faker/faker/v4/pkg/options"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"math/rand"
	"time"
)

var _ = Describe("Post", func() {
	var entityOne *Post
	var entityTwo *Post
	var err error

	textLength := uint(rand.Intn(TextMaxLength-TextMinLength+1) + TextMinLength)

	BeforeEach(func() {
		var text string

		_ = faker.FakeData(&text, options.WithRandomStringLength(textLength))
		entityOne, err = NewPost(faker.UUIDHyphenated(), text, nil, nil, nil)

		Expect(err).NotTo(HaveOccurred())

		_ = faker.FakeData(&text, options.WithRandomStringLength(textLength))
		timeNow := time.Now()
		entityTwo, err = NewPost(faker.UUIDHyphenated(), text, &timeNow, &timeNow, &timeNow)

		Expect(err).NotTo(HaveOccurred())
	})

	It("can transform an nil into nil", func() {
		findByIdQueryResponse, err := ToCqrsFindByIdQueryResponse(nil)

		Expect(err).NotTo(HaveOccurred())
		Expect(findByIdQueryResponse).To(BeNil())
	})

	It("can transform an entity into get_all_query_response", func() {
		entities := []Post{}
		entityPageableOne, err := NewEntityPageable(
			MinPageOrder,
			MinLimitItems,
			MinTotalItems,
		)
		entityPageableTwo, err := NewEntityPageable(
			rand.Intn(100)+MinPageOrder,
			MaxLimitItems,
			rand.Intn(100)+MinTotalItems,
		)

		Expect(err).NotTo(HaveOccurred())
		Expect(entityPageableOne).NotTo(BeNil())

		entities = []Post{*entityOne, *entityTwo}

		response, err := ToCqrsGetAllQueryResponse(&entities, entityPageableOne)

		Expect(err).NotTo(HaveOccurred())
		Expect(response).NotTo(BeNil())

		_, err = ToCqrsGetAllQueryResponse(&entities, entityPageableTwo)

		Expect(err).NotTo(HaveOccurred())

		Expect(response.Data[0].ID).To(Equal(entityOne.ID.GetId()))
		Expect(response.Data[0].Text).To(Equal(entityOne.Text.GetText()))
		Expect(response.Data[0].CreatedAt).To(BeNil())
		Expect(response.Data[0].UpdatedAt).To(BeNil())

		Expect(response.Data[1].ID).To(Equal(entityTwo.ID.GetId()))
		Expect(response.Data[1].Text).To(Equal(entityTwo.Text.GetText()))
		Expect(response.Data[1].CreatedAt).To(Equal(entityTwo.CreatedAt))
		Expect(response.Data[1].UpdatedAt).To(Equal(entityTwo.UpdatedAt))

		Expect(response.Pageable).To(Equal(entityPageableOne))
	})

	It("can transform an entity into find_by_id_query_response", func() {
		response, err := ToCqrsFindByIdQueryResponse(entityOne)

		Expect(err).NotTo(HaveOccurred())
		Expect(response).NotTo(BeNil())

		_, err = ToCqrsFindByIdQueryResponse(entityTwo)

		Expect(err).NotTo(HaveOccurred())

		Expect(response.Data.ID).To(Equal(entityOne.ID.GetId()))
		Expect(response.Data.Text).To(Equal(entityOne.Text.GetText()))
		Expect(response.Data.CreatedAt).To(BeNil())
		Expect(response.Data.UpdatedAt).To(BeNil())
		Expect(response.IsDeleted).To(BeFalse())
	})
})
