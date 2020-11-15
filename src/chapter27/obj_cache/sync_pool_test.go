package obj_cache

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
)

func TestSyncPool(t *testing.T) {
	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("Create a new Object")
			return 100
		},
	}

	v := pool.Get().(int)
	fmt.Println(v)
	pool.Put(3)
	// GC 会清除sync.pool中缓存的对象
	runtime.GC()
	v1, _ := pool.Get().(int)
	fmt.Println(v1)
}

func TestSyncMultiGroutine(t *testing.T) {
	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("Create a new Object")
			return 10
		},
	}

	pool.Put(100)
	pool.Put(200)
	pool.Put(300)

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			fmt.Printf("%d: %d\n", i, pool.Get())
			wg.Done()
		}(i)
	}
	wg.Wait()
}
