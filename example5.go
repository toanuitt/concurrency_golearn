package main
import "fmt"

func dup3(i <-chan int) (<- chan int,<- chan int,<- chan int){
	a,b,c := make(chan,int,2),make(chan,int,2),make(chan,int,2)
	go func(){
		x:= <- i
		a<-x
		b<-x
		c<-x
	}
	return a,b,c
}

func main() {
	x := fib()
	for i := 0; i < 10; i++ {
		fmt.Println(<-x)
	}
}