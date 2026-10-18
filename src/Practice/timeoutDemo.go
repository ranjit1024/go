package main
import (
	"fmt"
	"time"
)

func demo_timeout(){
	fmt.Println("Datasd");
	c1 := make(chan string, 1);
	go func(){
		time.Sleep()
	}
	
}