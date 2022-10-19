package util

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"testing"

	pool "github.com/jolestar/go-commons-pool/v2"
)

type myLocalConnection struct {
	name  string
	index int32
}

func TestSingleExample(t *testing.T) {

	v := int32(0)
	factory := pool.NewPooledObjectFactorySimple(func(ctx context.Context) (interface{}, error) {
		index := atomic.AddInt32(&v, 1)
		return &myLocalConnection{"actualConnect", index}, nil
	})

	ctx := context.Background()
	p := pool.NewObjectPoolWithDefaultConfig(ctx, factory)

	obj, err := p.BorrowObject(ctx)
	if err != nil {
		return
	}

	l := obj.(*myLocalConnection)
	fmt.Println(l.name, l.index)

	p.ReturnObject(ctx, obj)
}

func TestConcurrency(t *testing.T) {

	for i := 0; i < 10000; i++ {
		testValue := "yaoqijun"
		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			defer wg.Done()
			testValue = "yinsu"
		}()
		wg.Wait()
		if "yinsu" != testValue {
			t.Errorf("fail")
		}
	}

}
