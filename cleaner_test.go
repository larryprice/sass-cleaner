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
		styles, err := Parse(file)
		Expect(err).To(BeNil())
		testStyles := styles[".something"]
		Expect(testStyles).To(HaveLen(1))
		Expect(testStyles["margin-left"]).To(Equal("2em"))
	})
})
