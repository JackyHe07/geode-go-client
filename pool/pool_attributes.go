package pool

import (
	"fmt"
)


func init() {
	fmt.Println("pool attributes inited")
}


type LocationInfo struct {
	host string
	port int
}

type PoolAttributes struct {
	locations 	[]LocationInfo
}

func (pa *PoolAttributes) addLocator(host string, port int) {
	li := LocationInfo{host, port}
	pa.locations = append(pa.locations, li)
}

func (pa *PoolAttributes) getLocators() {
	for _, loc := range pa.locations {
		fmt.Println(loc)
	}
}
