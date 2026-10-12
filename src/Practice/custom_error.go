package main
import (
	"fmt"
	"errors"
)

type divError string{
	arg int
	message string
}
func(d *divError) Error() string{
	return fmt.Sprintf("%d  - %s", d.arg,d.message)
}
func div(a, b int)(int,error){
	if(a  == 0){
		return -1, &divError(a, "Cannot devide by 0")
	}
	return a / b, nil
}
func custom_error(){
	if ae, ok := errors
	fmt.Println("Hello World...")
}