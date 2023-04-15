package test

import (
	"fmt"
	"github.com/panjf2000/ants/v2"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

var wg sync.WaitGroup

func TestPool(t *testing.T) {

	// 创建一个协程池，最大协程数为10
	pool, err := ants.NewPool(4)
	if err != nil {
		panic(err)
	}
	defer pool.Release()

	// 定义一个计数器，用于记录所有任务的总执行次数
	var counter int64

	t1 := time.Now().Unix()
	// 添加100个任务到协程池中执行
	for i := 0; i < 10; i++ {
		wg.Add(1)
		pool.Submit(func() {
			// 执行任务
			atomic.AddInt64(&counter, 1)
			time.Sleep(time.Second * 1)
			defer wg.Done()
		})
	}

	// 等待所有任务执行完成
	wg.Wait()
	t2 := time.Now().Unix()

	fmt.Println("time:", t2-t1)
	// 输出执行结果
	fmt.Println("counter:", counter)
}
