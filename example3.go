package main

import (
	"fmt"
	"runtime"
	"time"
)

func printNumbers() {
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
		time.Sleep(500 * time.Millisecond)
	}
}

func printLetters() {
	for i := 'A'; i <= 'E'; i++ {
		fmt.Println(string(i))
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	runtime.GOMAXPROCS(2) // Cho phép 2 CPU chạy đồng thời

	go printNumbers()
	go printLetters()

	time.Sleep(3 * time.Second) // Đợi các goroutine kết thúc
}
