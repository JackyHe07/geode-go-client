package cache 

import (
	"fmt"
)

type Cache struct {
	dsProp  	Properties
}


func (c *Cache) GetProp() {
	for k, v := range c.dsProp {
		fmt.Println(k, v)
	}
}
