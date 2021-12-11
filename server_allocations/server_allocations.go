package server_allocations

import (
	"fmt"
	"sort"

	"github.com/nawazish-github/consistent-hashing/hash"
)

type ServerAllocation struct {
	serverLocations map[int]string
}

func (sa *ServerAllocation) InitServerAllocation() {
	sa.serverLocations = make(map[int]string)
}

func (sa *ServerAllocation) AllocateServer(serverKey string) int {
	loc := hash.LocationOnRing(serverKey)
	sa.serverLocations[loc] = serverKey
	return loc
}

func (sa *ServerAllocation) FindTheServer(reqKey string) {
	loc := hash.LocationOnRing(reqKey)
	if server, ok := sa.serverLocations[loc]; ok {
		fmt.Printf("Request %s is served by server %s", reqKey, server)
	} else {
		sa.walk(loc, reqKey)
	}
}

func (sa *ServerAllocation) walk(loc int, reqKey string) string {
	sortedKeys := make([]int, len(sa.serverLocations))

	for k, _ := range sa.serverLocations {
		sortedKeys = append(sortedKeys, k)
		sortedKeys = sort.IntSlice(sortedKeys)
	}

	for i, serLoc := range sortedKeys {
		if serLoc > loc {
			fmt.Printf("Request %s would be served on Server %s", sa.serverLocations[i], reqKey)
			return sa.serverLocations[i-1]
		}
	}
	panic(`no server available to serve request #{reqKey}`)
}
