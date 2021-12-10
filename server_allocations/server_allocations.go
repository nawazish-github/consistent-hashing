package server_allocations

import "github.com/nawazish-github/consistent-hashing/hash"

type ServerAllocation struct {
	serverLocations map[int64]string
}

func (sa *ServerAllocation) InitServerAllocation() {
	sa.serverLocations = make(map[int64]string)
}

func (sa *ServerAllocation) AllocateServer(serverKey string) {
	loc := hash.LocationOnRing(serverKey)
	sa.serverLocations[loc] = serverKey
}

func (sa *ServerAllocation) FindTheServer(reqKey string) {

}
