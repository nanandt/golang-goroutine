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

func OnlyIn(channel chan<- string){ // hanya boleh mengirim data ke channel
	time.Sleep(2 *time.Second)
	channel <- "Bakul Sega"
}

func OnlyOut(channel <-chan string){ // hanya boleh mengambil data dari channel
	data := <-channel
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(3 *time.Second)

}

func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 3)  // 3 slot channel
	defer close(channel)

	go func() {
		channel <- "Jaki"
		channel <- "Maulana"
		channel <- "Fikri"
	}()

	go func() {
		fmt.Println(<- channel)
		fmt.Println(<- channel)
		fmt.Println(<- channel)
	}()

	time.Sleep(2 *time.Second)
	fmt.Println("Selesai")
}

