package golang_goroutine

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestGetGoMaxprocs(t *testing.T) {
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++{
		wg.Add(1)
		go func() {
			time.Sleep(3 *time.Second)
			wg.Done()
		}()
	}
	totalcpu := runtime.NumCPU()
	fmt.Println("CPU = ", totalcpu)

	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("total thread = ", totalThread)

	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("total goroutine = ", totalGoroutine)

	wg.Wait()
}

