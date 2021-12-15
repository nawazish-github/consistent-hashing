package server_allocations

import (
	"fmt"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestCart(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Shopping Cart Suite")
}

var _ = Describe("Server tests", func() {
	var serverAlloc ServerAllocation
	BeforeEach(func() {
		serverAlloc = ServerAllocation{}
		serverAlloc.InitServerAllocation()
	})
	Context("Attempting to allocate server on the hash ring", func() {
		When("A server key is given", func() {
			It("Should allocate on the hash ring", func() {
				serverKey := "198.8.9.3"
				loc := serverAlloc.AllocateServer(serverKey)
				Expect(serverAlloc.serverLocations[loc]).Should(Equal("198.8.9.3"))
			})
		})
	})

	Context("Attempt to find a server", func() {
		When("A request key is provided and at max one server is available on the cluster", func() {
			It("Should route the request to the available server", func() {
				serverKey := "com.server"
				serverAlloc.AllocateServer(serverKey)
				requestKey := "get me coffee"
				server := serverAlloc.FindTheServer(requestKey)
				fmt.Println("SERVER BABA: ", server)
				Expect(server).Should(Equal("com.server"))
			})
		})
	})
})
