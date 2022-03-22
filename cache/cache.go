package cache 

import (
	"fmt"
	"github.com/jackyhe07/geode-go-client/pool"
)

type Cache struct {
	dsProp  	Properties
	poolManager	*pool.PoolManager
}


func (c *Cache) init() {
	c.poolManager = &pool.PoolManager{}
	c.poolManager.Init()
	fmt.Println("cache inited")
}

func (c *Cache) GetProp() {
	for k, v := range c.dsProp {
		fmt.Println(k, v)
	}
}

func (c *Cache) GetPoolManager() *pool.PoolManager {
	return c.poolManager
}
