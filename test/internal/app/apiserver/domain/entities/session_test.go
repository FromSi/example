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

var _ = Describe("Session", func() {
	agentLength := uint(rand.Intn(AgentSessionMaxLength-AgentSessionMinLength+1) + AgentSessionMinLength)

	It("can make a session", func() {
		var agentOne string
		var agentTwo string

		_ = faker.FakeData(&agentOne, options.WithRandomStringLength(agentLength))
		timeNowOne := time.Now()
		ipOne := faker.IPv4()
		session, err := NewSession("", "", ipOne, agentOne, &timeNowOne, &timeNowOne, &timeNowOne, &timeNowOne, &timeNowOne)

		Expect(err).NotTo(HaveOccurred())
		Expect(session).NotTo(BeNil())

		_ = faker.FakeData(&agentTwo, options.WithRandomStringLength(agentLength))
		timeNowTwo := time.Now()
		ipTwo := faker.IPv6()
		_, err = NewSession("", "", ipTwo, agentTwo, &timeNowTwo, &timeNowTwo, &timeNowTwo, &timeNowTwo, &timeNowTwo)

		Expect(err).NotTo(HaveOccurred())

		Expect(session.IP.GetIp()).To(Equal(ipOne))
		Expect(session.Agent.GetAgent()).To(Equal(agentOne))
		Expect(session.CreatedAt).To(Equal(&timeNowOne))
		Expect(session.UpdatedAt).To(Equal(&timeNowOne))
		Expect(session.DeletedAt).To(Equal(&timeNowOne))
	})

	It("can decline a session", func() {
		var agentOne string

		_ = faker.FakeData(&agentOne, options.WithRandomStringLength(agentLength))
		ipOne := faker.IPv4()

		session, err := NewSession("", "", ipOne, agentOne, nil, nil, nil, nil, nil)

		Expect(err).NotTo(HaveOccurred())
		Expect(session).NotTo(BeNil())

		Expect(session.DeclinedAt).To(BeNil())

		err = session.Decline(nil)

		Expect(err).NotTo(HaveOccurred())

		Expect(session.DeclinedAt).NotTo(BeNil())

		timeNowOne := time.Now()

		err = session.Decline(&timeNowOne)

		Expect(err).NotTo(HaveOccurred())

		Expect(*session.DeclinedAt).To(Equal(timeNowOne))
	})

	Describe("ID Value Object", func() {
		var idValueObject *IdSession
		var err error

		It("can make correctly an id", func() {
			idValueObject, err = NewIdSession("")

			Expect(err).NotTo(HaveOccurred())
			Expect(idValueObject).NotTo(BeNil())

			idValueObject, err = NewIdSession(faker.Word())

			Expect(err).To(HaveOccurred())
			Expect(idValueObject).To(BeNil())

			idOne := faker.UUIDHyphenated()
			idValueObject, err = NewIdSession(idOne)

			Expect(err).NotTo(HaveOccurred())
			Expect(idValueObject).NotTo(BeNil())

			idTwo := faker.UUIDHyphenated()
			_, err = NewIdSession(idTwo)

			Expect(err).NotTo(HaveOccurred())

			Expect(idValueObject.GetId()).To(Equal(idOne))
		})
	})

	Describe("IP Value Object", func() {
		var ipValueObject *IpSession
		var err error

		It("can make correctly an ip", func() {
			ipValueObject, err = NewIpSession("")

			Expect(err).To(HaveOccurred())
			Expect(ipValueObject).To(BeNil())

			ipValueObject, err = NewIpSession(faker.Word())

			Expect(err).To(HaveOccurred())
			Expect(ipValueObject).To(BeNil())

			ipOne := faker.IPv4()
			ipValueObject, err = NewIpSession(ipOne)

			Expect(err).NotTo(HaveOccurred())
			Expect(ipValueObject).NotTo(BeNil())

			ipTwo := faker.IPv6()
			_, err = NewIpSession(ipTwo)

			Expect(err).NotTo(HaveOccurred())

			Expect(ipValueObject.GetIp()).To(Equal(ipOne))
		})
	})

	Describe("Agent Value Object", func() {
		var agentValueObject *AgentSession
		var agentOne string
		var agentTwo string
		var err error

		It("can make correctly an agent", func() {
			agentValueObject, err = NewAgentSession("")

			Expect(err).To(HaveOccurred())
			Expect(agentValueObject).To(BeNil())

			_ = faker.FakeData(&agentOne, options.WithRandomStringLength(AgentSessionMinLength-1))
			agentValueObject, err = NewAgentSession(agentOne)

			Expect(err).To(HaveOccurred())
			Expect(agentValueObject).To(BeNil())

			_ = faker.FakeData(&agentOne, options.WithRandomStringLength(AgentSessionMaxLength+1))
			agentValueObject, err = NewAgentSession(agentOne)

			Expect(err).To(HaveOccurred())
			Expect(agentValueObject).To(BeNil())

			_ = faker.FakeData(&agentOne, options.WithRandomStringLength(agentLength))
			agentValueObject, err = NewAgentSession(agentOne)

			Expect(err).NotTo(HaveOccurred())
			Expect(agentValueObject).NotTo(BeNil())

			_ = faker.FakeData(&agentTwo, options.WithRandomStringLength(agentLength))
			_, err = NewAgentSession(agentTwo)

			Expect(err).NotTo(HaveOccurred())

			Expect(agentValueObject.GetAgent()).To(Equal(agentOne))
		})
	})

	Describe("Refresh Token Value Object", func() {
		var refreshTokenValueObjectOne *RefreshTokenSession
		var refreshTokenValueObjectTwo *RefreshTokenSession
		var err error

		It("can make correctly a refresh_token", func() {
			refreshTokenValueObjectOne, err = NewRefreshTokenSession()

			Expect(err).NotTo(HaveOccurred())
			Expect(refreshTokenValueObjectOne).NotTo(BeNil())

			refreshTokenValueObjectTwo, err = NewRefreshTokenSession()

			Expect(err).NotTo(HaveOccurred())
			Expect(refreshTokenValueObjectTwo).NotTo(BeNil())

			Expect(refreshTokenValueObjectOne.GetToken()).NotTo(Equal(refreshTokenValueObjectTwo.GetToken()))
		})
	})
})
