package main

import "fmt"

func main() {
	ch := make(chan int) // Unbuffered channel

	fmt.Println("Mengirim data ke channel...")
	ch <- 42                     // Terblokir karena tidak ada penerima
	fmt.Println("Data terkirim") // Tidak akan dieksekusi
}
