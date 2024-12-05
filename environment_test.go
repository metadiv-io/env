package env

import (
	"os"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Environment", func() {
	Describe("New", func() {
		It("should return an error if the environment variable is not set and required is true", func() {
			Expect(func() { New("NON_EXISTENT_VARIABLE", true) }).To(Panic())
		})
		It("should return an environment variable if the environment variable is set", func() {
			os.Setenv("EXISTENT_VARIABLE", "value")
			Expect(New("EXISTENT_VARIABLE", true).String()).To(Equal("value"))
		})
		It("should return an environment variable if the environment variable is set and required is false", func() {
			os.Setenv("EXISTENT_VARIABLE", "value")
			Expect(New("EXISTENT_VARIABLE", false).String()).To(Equal("value"))
		})
	})

	Describe("String", func() {
		It("should return the environment variable value as a string", func() {
			os.Setenv("EXISTENT_VARIABLE", "value")
			Expect(New("EXISTENT_VARIABLE", true).String()).To(Equal("value"))
		})
	})

	Describe("Strings", func() {
		It("should return the environment variable value as a slice of strings", func() {
			os.Setenv("EXISTENT_VARIABLE", "value1,value2")
			Expect(New("EXISTENT_VARIABLE", true).Strings()).To(Equal([]string{"value1", "value2"}))
		})
	})

	Describe("Int", func() {
		It("should return the environment variable value as an integer", func() {
			os.Setenv("EXISTENT_VARIABLE", "123")
			Expect(New("EXISTENT_VARIABLE", true).Int()).To(Equal(123))
		})
	})

	Describe("Uint", func() {
		It("should return the environment variable value as an unsigned integer", func() {
			os.Setenv("EXISTENT_VARIABLE", "123")
			Expect(New("EXISTENT_VARIABLE", true).Uint()).To(Equal(uint(123)))
		})
	})

	Describe("Float", func() {
		It("should return the environment variable value as a float", func() {
			os.Setenv("EXISTENT_VARIABLE", "123.45")
			Expect(New("EXISTENT_VARIABLE", true).Float()).To(Equal(123.45))
		})
	})

	Describe("Bool", func() {
		It("should return the environment variable value as a boolean", func() {
			os.Setenv("EXISTENT_VARIABLE", "true")
			Expect(New("EXISTENT_VARIABLE", true).Bool()).To(BeTrue())
		})
	})

	AfterEach(func() {
		os.Clearenv()
	})
})
