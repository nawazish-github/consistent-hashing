package server_allocations

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Test server allocation", func() {
	var serverAlloc ServerAllocation
	BeforeEach(func() {
		serverAlloc = ServerAllocation{}
		serverAlloc.InitServerAllocation()
	})
	When("A server key is given", func() {
		It("Should allocate on the hash ring", func() {
			serverKey := "198.8.9.3"
			loc := serverAlloc.AllocateServer(serverKey)
			Expect(serverAlloc.serverLocations[loc]).Should(Equal("198.8.9.3"))
		})
	})
})
