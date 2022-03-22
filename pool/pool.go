package pool

import (
	"fmt"
)


type Pool struct {
	pa 	PoolAttributes
}

func init() {
	fmt.Println("pool inited")
}
