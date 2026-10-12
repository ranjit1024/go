package main
import (
	"fmt"
	"errors"
)

func divide(a , b int) (int,error){
	if(a == 0){
		return -1, errors.New("Can not devide by 0")
	}
	return a / b , nil;
}
func error_demo(){
	res , err := divide(12,12);
	fmt.Println(res)
	fmt.Println(err)
	fmt.Println("Hellow World...")
}