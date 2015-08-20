package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestSassCleaner(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "SassCleaner Suite")
}
