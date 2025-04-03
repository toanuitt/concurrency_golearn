package main

import (
	"fmt"
	"time"
)

func ready(w string, sec int){
	time.Sleep(time.Duration(sec)*time.Second)
	fmt.Printf(w,"is ready")
}

func main(){
	go ready("Tea",2)
	go ready("Cofee",1)
	fmt.Printf("Im waiting")
	time.Sleep(5*time.Second)
}