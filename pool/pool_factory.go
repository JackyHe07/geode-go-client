package pool

import (
	"fmt"
)


type PoolFactory struct {
	pa	PoolAttributes
}

func init() {
	fmt.Println("pool factory inited")
}

func (pf *PoolFactory) init() {
}

func (pf *PoolFactory) AddLocator(host string, port int) *PoolFactory{
	pf.pa.addLocator(host, port)
	pf.pa.getLocators()
	return pf
}

func (pf *PoolFactory) Create() *Pool {
	pool := &Pool{pf.pa}
	return pool
}
