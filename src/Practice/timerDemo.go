package main
import (
	"fmt"
	"time"
)

func timer_demo(){
	timer1 := time.NewTimer(2 * time.Second);

	<- timer1.C;
	fmt.Println("Timer 1 fired")
	fmt.Println("Hello World..");
}