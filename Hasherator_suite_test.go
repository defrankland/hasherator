package Hasherator_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestHasherator(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Hasherator Suite")
}
