package main 

import (
	"fmt"
	"github.com/jackyhe07/geode-go-client/cache"
)

func main() {
	fmt.Println("testing ...")
	cache := cache.GetCacheFactory().
			Set("name", "test").
			Set("log-level","info").
			Create()
	cache.GetProp()
	cf := cache.GetPoolManager().CreateFactory()
	cf.AddLocator("localhost", 12345).AddLocator("127.0.0.1", 32135).Create()
}
