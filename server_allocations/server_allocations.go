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

func (sa *ServerAllocation) FindTheServer(reqKey string) string {
	loc := hash.LocationOnRing(reqKey)
	if server, ok := sa.serverLocations[loc]; ok {
		fmt.Printf("Request %s is served by server %s", reqKey, server)
		return server
	} else {
		return sa.walk(loc, reqKey)
	}
}

func (sa *ServerAllocation) walk(loc int, reqKey string) string {
	if len(sa.serverLocations) == 0 {
		panic("no server available on cluster to process request")
	}

	var sortedKeys []int

	for k := range sa.serverLocations {
		sortedKeys = append(sortedKeys, k)
	}

	sortedKeys = sort.IntSlice(sortedKeys)

	for i := loc; ; {
		for j := 0; j < len(sortedKeys); j++ {
			if i <= sortedKeys[j] {
				fmt.Printf("Request %s would be served on server %s", reqKey, sa.serverLocations[sortedKeys[j]])
				return sa.serverLocations[i]
			}
		}

		//reset the loop counter as soon it exhausts the output space
		//so that a server is searched on locations behind hashed value
		//of the request on the hash ring
		if i == hash.OutputSpace {
			i = 0
			continue
		}

		//if the loop counter has searched the entire output space
		//and yet no server was assigned to the request, it would
		//mean there is no server on the cluster, panic out
		if i == loc-1 {
			panic(`no server available to serve request #{reqKey}`)
		}

		i++
	}

}
