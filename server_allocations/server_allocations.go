package server_allocations

import (
	"fmt"

	"github.com/nawazish-github/consistent-hashing/hash"
)

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
	loc := hash.LocationOnRing(reqKey)
	if server, ok := sa.serverLocations[loc]; ok {
		fmt.Printf("Request %s is served by server %s", reqKey, server)
	} else {
		sa.walk(loc, reqKey)
	}
}

func (sa *ServerAllocation) walk(loc int64, reqKey string) string {

	//15, 30, 45, 60, 67, 88, 90 | 58

	sortedKeys := make([]int64, len(sa.serverLocations))

	for k, _ := range sa.serverLocations {
		sortedKeys = append(sortedKeys, k)
		fmt.Println(sortedKeys)
	}
	return ""
}
