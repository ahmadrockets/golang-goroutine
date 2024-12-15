---
title: "Golang Goroutine"
description: "This post will explain golang goroutine and channels and also wait group"
publishDate: "01 Nov 2024"
updatedDate: "01 Nov 2024"
# coverImage:
#   src: "./cover.png"
#   alt: "Astro build wallpaper"
tags: ["golang", "tech", "goroutine", "teknologi"]
---

# GOROUTINE & CHANNELS
## Goroutine
Goroutine adalah lightweight thread yang dimanage oleh Go runtime. Goroutine sangat ringan, hanya dibutuhkan sekitar **2kB** memori saja untuk satu buah Goroutine. Goroutine memiliki sifat yang asynchronous jadi tidak saling menunggu dengan Goroutine yang lain. 

Untuk membuat Goroutine baru caranya cukup mudah, yaitu dengan menambahkan awalan `go` yang diikuti dengan nama method yang akan dijalankan secara Goroutine. Berikut ini contoh penggunaannya.
```go
package main

import (
  "fmt"
  "time"
)

func say(s string) {
  for i := 0; i < 5; i++ {
   time.Sleep(100 * time.Millisecond)
   fmt.Println(s)
  }
}

func main() {
  go say("world")
  say("hello")
}
```
Jika program dijalankan maka akan memunculkan output seperti berikut :
```bash
fahrudin@belajar-goroutine $ go run main.go
world
hello
hello
world
world
hello
hello
world
world
hello    
```
Dapat dilihat output yang dihasilkan adalah tulisan `hello` dan `world` muncul secara selang seling, ini dikarenakan `say("world")` dijalankan sebagai Goroutine, sehingga tidak terjadi saling menunggu.

## Channels
Channel adalah penghubung antara goroutine satu dengan goroutine lainnya. Channel bisa dibuat dengan menggunakan fungsi `make()` dengan menentukan tipe data yang akan dikirim melalui channel.

``` go
ch := make(chan int)
```
kode diatas adalah contoh kode untuk membuat channel.

``` go
ch <- v    // Mengirim variable v ke channel ch.
v := <-ch  // Menerima dari ch ch, dan assign value ke v
```
Kode diatas adalah contoh cara untuk mengirim dan menerima channel.

```go
package main

import "fmt"

func cetak(ch chan int, angka int) {
	fmt.Println("ini dari goroutine cetak...")
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
```
Kode diatas adalah contoh penggunaan channel, jika dijalankan 2x maka akan menghasilkan output:
```bash
fahrudin@belajar-goroutine $ go run channel/channel.go
ini dari goroutine cetak... 30
ini dari goroutine cetak... 20
nilai channel integer 1 : 30
nilai channel integer 2 : 20
ini dari function main...

fahrudin@belajar-goroutine $ go run channel/channel.go
ini dari goroutine cetak... 20
ini dari goroutine cetak... 30
nilai channel integer 1 : 20
nilai channel integer 2 : 30
ini dari function main...
```
Output pertama dan kedua berbeda urutan tergantung goroutine mana yang lebih dulu di eksekusi. Hal ini dikarenakan, pengiriman data adalah dari 2 goroutine yang berbeda, yang kita tidak tau mana yang prosesnya selesai lebih dulu. Goroutine yang dieksekusi lebih awal belum tentu selesai lebih awal, yang jelas proses yang selesai lebih awal datanya akan diterima lebih awal.

## Blocking Channels
Pengiriman dan penerimaan data pada channel bersifat `blocking` atau synchronous. Artinya, statement di-bawah syntax pengiriman dan penerimaan data via channel hanya akan dieksekusi setelah proses serah terima berlangsung dan selesai.
```go
package main

import "fmt"

func main() {
	ch := make(chan int) // Unbuffered channel

	fmt.Println("Mengirim data ke channel...")
	ch <- 42                     // Terblokir karena tidak ada penerima
	fmt.Println("Data terkirim") // Tidak akan dieksekusi
}

```
Contoh kode diatas jika dijalankan maka baris kode terakhir tidak akan pernah di eksekusi, akan terjadi proses deadlock dan program akan excited ```fatal error: all goroutines are asleep - deadlock!``` karena tidak ada proses penerimaan channel dari proses diatas.


## Buffered Channels
Secara default channel bersifat *unbufferd*. buffered channel sama seperti channel biasa, tetapi buffered channel memiliki kapasitas penyimpanan data. Selama jumlah data yang dikirim tidak melebihi jumlah buffer, maka pengiriman akan berjalan ***asynchronous*** (tidak blocking). Untuk membuat buffered channel kita bisa menggunakan fungsi ```make(chan string, 2)```, artinya channel dibuat dengan kapasitas 2. <br />
Kelabihan dari buffered channel antara lain bisa mengurangi blocking pada goroutine, selain itu data bisa ditampung sementara tanpa menunggu goroutine penerima. Untuk kekurangan dari buffered channel adalah jika channel penuh, goroutine akan diblock sampai ada tempat kosong.

```go
package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
```
Kode diatas artinya kita membuat channel dengan kapasitas 2 buffers. Jika dijalankan akan menampilkan.
```bash
fahrudin@belajar-goroutine $ go run buffered_channel/buffered_channel.go 
1
2
```

Berikut ini contoh code channel lainnya.
```go
package main

import "fmt"

func main() {
	ch2 := make(chan int, 2)
	ch2 <- 1
	ch2 <- 2
	ch2 <- 3
	fmt.Println(<-ch2)
	fmt.Println(<-ch2)
	fmt.Println(<-ch2)
}
```

Jika dijalankan maka akan terjadi error deadlock karena kapasitas channel yang didefinisikan adalah 2 buffers, tetapi ada 3 kali value dikirimkan ke channel tersebut.
```bash
fahrudin@belajar-goroutine $ go run buffered_channel/buffered_channel.go
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan send]:
main.main()
        /golang-goroutine/buffered_channel/buffered_channel.go:9 +0x58
exit status 2
```

## Channel Select
## Channel Range & Close
## Channel Timeout
## WaitGroup
## Penerapan Goroutine


## Referensi Tulisan
- https://dasarpemrogramangolang.novalagung.com/A-goroutine.html
- https://dasarpemrogramangolang.novalagung.com/A-channel.html
- https://medium.com/@jamal.kaksouri/goroutines-in-golang-understanding-and-implementing-concurrent-programming-in-go-600187bcfaa2
- https://buildwithangga.com/tips/penggunaan-channel-dalam-go-komunikasi-antar-go-routines
