package entities_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"testing"
)

func TestEntities(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Domain Entities")
}
