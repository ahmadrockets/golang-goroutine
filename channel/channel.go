package main

import "fmt"

func cetak(ch chan int, angka int) {
	fmt.Println("ini dari goroutine cetak...", angka)
	ch <- angka
}

func main() {
	angka := make(chan int)
	go cetak(angka, 20)
	go cetak(angka, 30)
	nilai1, nilai2 := <-angka, <-angka
	fmt.Println("nilai channel integer 1 :", nilai1)
	fmt.Println("nilai channel integer 2 :", nilai2)
	fmt.Println("ini dari function main...")
}
