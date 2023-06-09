// go build sleepsort.go && ./sleepsort
package sort

import (
	"sync"
	"time"
)

func Sleepsort(arr []int) []int {
	var wg sync.WaitGroup
	out := make(chan int, len(arr))

	for _, v := range arr {
		wg.Add(1)
		go func(val int) {
			defer wg.Done()
			time.Sleep(time.Duration(val) * time.Second)
			out <- val
		}(v)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	result := make([]int, 0, len(arr))
	for v := range out {
		result = append(result, v)
	}
	return result
}
