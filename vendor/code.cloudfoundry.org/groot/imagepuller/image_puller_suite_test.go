package imagepuller_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestImagePuller(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Imagepuller suite")
}
