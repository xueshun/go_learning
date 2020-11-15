package obj_pool

import (
	"fmt"
	"testing"
	"time"
	"unsafe"
)

func TestObjPool(t *testing.T) {
	pool := NewObjPool(10)
	// 尝试放置超出池大小的对象
	/*if err := pool.ReleaseObj(&ReusableObj{}); err != nil {
		t.Error(err)
	}*/
	for i := 0; i < 11; i++ {
		if v, err := pool.GetObj(time.Second * 1); err != nil {
			t.Error(err)
		} else {
			fmt.Printf("%d %T\n %v\n", i, v, unsafe.Pointer(v))
			if err := pool.ReleaseObj(v); err != nil {
				t.Error(err)
			}
		}
	}
}
