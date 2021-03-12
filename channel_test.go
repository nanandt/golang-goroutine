package golang_goroutine

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string) // channel yg menerima tipe data string
	//channel <- "Rizky" // mengirim data ke channel
	//data := <- channel // mengambil data ke channel
	//fmt.Println(<- channel) // mengambil data channel dan dimasukan ke parameter
	defer close(channel) // selesai membuat channel direkomendasikan menutup channel nya

	go func() {
		time.Sleep(2 *time.Second)
		channel <- "Muhamad Fatikhurrizky"
		fmt.Println("Selesai mengirim data ke channel")
	}()

	data := <-channel //mengambil data dari channel
	fmt.Println(data)
	time.Sleep(5 *time.Second)
}

func GiveMeResponse(channel chan string){
	time.Sleep(2 *time.Second)
	channel <- "Harun Bakul Gas"
	fmt.Println("Selesai mengirim data ke channel")
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go GiveMeResponse(channel)

	data := <-channel
	fmt.Println(data)

	time.Sleep(5 *time.Second)
}


