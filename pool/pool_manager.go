package pool

import (
	"fmt"
)


type PoolManager struct {
}

func (pm *PoolManager) Init() {
	fmt.Println("pool manager inited")
}


func (pm *PoolManager) CreateFactory() *PoolFactory {
	pf := &PoolFactory{}
	pf.init()
	return pf
}
