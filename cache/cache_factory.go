package cache 

import (
	"fmt"    
	//"strings"
)


type CacheFactory struct {
	dsProp 	 Properties 
}

var cf *CacheFactory


func init() {
	cf = &CacheFactory{}
	cf.init()
	fmt.Println("cache factory inited")
}

func GetCacheFactory() *CacheFactory {
	return cf
}


func (cf *CacheFactory) init() {
	cf.dsProp = make(Properties)
}

func (cf *CacheFactory) Set(name, value string) *CacheFactory{
	cf.dsProp[name] = value
	return cf
}

func (cf *CacheFactory) Create() Cache {
	cache := Cache{dsProp:cf.dsProp}
	cache.init()
	return cache
}
