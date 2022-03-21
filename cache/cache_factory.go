package cache 

import (
	//"fmt"    
	//"strings"
)


type CacheFactory struct {
	dsProp 	 Properties 
}

var cf *CacheFactory


func GetCacheFactory() *CacheFactory {
	if cf == nil {
		cf = &CacheFactory{}
		cf.init()
	}
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
	return cache
}
