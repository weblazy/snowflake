package snowflake

import (
	"fmt"
	"testing"
)

func TestSnowFlake(t *testing.T) {
	// 测试脚本

	// 生成节点实例
	worker, err := NewWorker(1)

	if err != nil {
		fmt.Println(err)
		return
	}

	ch := make(chan int64)
	count := 400000
	// 并发 count 个 goroutine 进行 snowflake id 生成
	for i := 0; i < count; i++ {
		go func() {
			id := worker.GetId()
			ch <- id
		}()
	}

	defer close(ch)

	m := make(map[int64]int)
	for i := 0; i < count; i++ {
		id := <-ch
		// 如果 map 中存在为 id 的 key, 说明生成的 snowflake id 有重复
		_, ok := m[id]
		if ok {
			t.Error(id, "出现重复id!\n")
			return
		}
		// 将 id 作为 key 存入 map
		m[id] = i
	}
	// 成功生成 snowflake id
	fmt.Printf("总共生成%d个id\n", count)
}
