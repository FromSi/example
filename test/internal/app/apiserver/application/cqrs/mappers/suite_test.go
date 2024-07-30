package mappers_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"testing"
)

func TestApplicationCQRSMappers(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Application CQRS Mappers")
}
