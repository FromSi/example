package tools

import (
	. "github.com/fromsi/example/internal/pkg/tools"
	"github.com/go-faker/faker/v4"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"time"
)

const ExpirationInMinutes = 30

var _ = Describe("JWT", func() {
	secretKey := "test"

	Describe("Session JWT", func() {
		It("can make correctly jwt and decode", func() {
			issuer := faker.Word()
			audience := faker.Word()
			subject := faker.Word()
			timeNow := time.Now()

			sessionJWT := NewSessionJWT(issuer, audience, subject, timeNow, ExpirationInMinutes)

			jwtToken, err := sessionJWT.GetJWT(secretKey)

			Expect(err).To(BeNil())

			sessionJWT, err = NewSessionJWTFromString(jwtToken, secretKey)

			Expect(err).To(BeNil())

			Expect(sessionJWT.Issuer).To(Equal(issuer))
			Expect(sessionJWT.Audience).To(Equal(audience))
			Expect(sessionJWT.Subject).To(Equal(subject))
		})

		It("can check the issued at", func() {
			issuer := faker.Word()
			audience := faker.Word()
			subject := faker.Word()
			timeNow := time.Now().AddDate(0, 0, 1)

			sessionJWT := NewSessionJWT(issuer, audience, subject, timeNow, ExpirationInMinutes)

			tokenJWT, err := sessionJWT.GetJWT(secretKey)

			Expect(err).NotTo(HaveOccurred())

			_, err = NewSessionJWTFromString(tokenJWT, secretKey)

			Expect(err).To(HaveOccurred())
		})

		It("can check the expiration date", func() {
			issuer := faker.Word()
			audience := faker.Word()
			subject := faker.Word()
			timeNow := time.Now().Add(-time.Minute * ExpirationInMinutes)

			sessionJWT := NewSessionJWT(issuer, audience, subject, timeNow, ExpirationInMinutes)

			tokenJWT, err := sessionJWT.GetJWT(secretKey)

			Expect(err).NotTo(HaveOccurred())

			_, err = NewSessionJWTFromString(tokenJWT, secretKey)

			Expect(err).To(HaveOccurred())
		})
	})
})
