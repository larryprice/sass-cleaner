package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"os"
)

var _ = Describe("Cleaner", func() {
	var file *os.File
	BeforeSuite(func () {
		var err error
	  file, err = os.Open("test/test.scss");
		Expect(err).To(BeNil())
	})

	AfterSuite(func() {
		defer file.Close()
	})

	It("Parses simple sass file", func () {
		Expect(Parse(file)).To(BeNil())
		// Expect(true).To(Equal(false))
	})
})
