package repositories_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"testing"
)

func TestInfrastructureRepositories(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Infrastructure Repositories")
}
