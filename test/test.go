package main 

import (
	"fmt"
	"github.com/jackyhe07/geode-go-client/cache"
)

func main() {
	fmt.Println("testing ...")
	cache := cache.GetCacheFactory().Set("name", "test").Set("log-level","info").Create()
	cache.GetProp()
}
