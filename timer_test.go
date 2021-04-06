package golang_goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	timer := time.NewTimer(5 *time.Second)
	fmt.Println(time.Now())

	time := <- timer.C
	fmt.Println(time)
}

func TestAfter(t *testing.T) {
	channel := time.After(5 *time.Second)
	fmt.Println(time.Now())

	time := <- channel
	fmt.Println(time)
}

func TestAfterFunc(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(1)
	time.AfterFunc(5 *time.Second, func() {
		fmt.Println(time.Now())
		wg.Done()
	})
	fmt.Println(time.Now())
	wg.Wait()
}


