package golang_goroutine

import (
	"fmt"
	"testing"
	"time"
)

func RunHelloWorld(){
	fmt.Println("Hello World")
}

// goroutine tdk cocok jika ada func yg ada return value nya


func TestCreateGoroutine(t *testing.T){
	go RunHelloWorld()
	fmt.Println("Ups")

	time.Sleep(1 *time.Second)

}