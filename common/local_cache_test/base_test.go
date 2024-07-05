package local_cache_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"git.code.oa.com/trpc-go/trpc-database/localcache"
)

func TestL(t *testing.T) {
	fmt.Println("===============")
	ctx := context.Background()
	cache := localcache.New(localcache.WithCapacity(64), localcache.WithExpiration(60*60))
	value, _ := cache.GetWithCustomLoad(ctx, "haha", func(ctx context.Context, key string) (interface{}, error) {
		fmt.Println("set haha: haha_haha")
		return "haha_haha", nil
	}, 3)
	fmt.Println("get haha: ", value)
	value, flag := cache.Get("haha")
	fmt.Println("after zero second, get haha: ", value, flag, cache.Len())
	time.Sleep(time.Second)
	value, flag = cache.Get("haha")
	fmt.Println("after one second, get haha: ", value, flag, cache.Len())
	time.Sleep(time.Second)
	value, flag = cache.Get("haha")
	fmt.Println("after two second, get haha: ", value, flag, cache.Len())
	time.Sleep(time.Second)
	value, flag = cache.Get("haha")
	fmt.Println("after three second, get haha: ", value, flag, cache.Len())
	time.Sleep(time.Second)
	value, flag = cache.Get("haha")
	fmt.Println("after four second, get haha: ", value, flag, cache.Len())
}
