package main 

import ("fmt" 
		"time"
)

func main(){

	go doSomething();
	doSomething();
	
	fmt.Println("Hello go")
	go func(){
		fmt.Println("Hello 123")
	}()
	go func(){
		fmt.Println("Hello om")
	}()
	
	fmt.Println("Main end")
	
	time.Sleep(1 * time.Second)

	
	
}

func doSomething(){
	fmt.Println("Did Something...")
}