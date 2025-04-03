package main
import (
	"fmt"
	"time"
)

var c chan int
func ready(w string, sec int){
	time.Sleep(time.Duration(sec)*time.Second)
	fmt.Println(w,"is ready")
	c <-1
}
func main(){
	c = make(chan int)
	go ready("tea",2)
	go ready("coffee",1)
	fmt.Println("Im waiting for but too long")
	<-c
	<-c	
}