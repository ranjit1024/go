package main
import (
	"fmt"
	"errors"
)

type divError struct{
	arg int
	message string
}
func(d *divError) Error() string{
	return fmt.Sprintf("%d  - %s", d.arg,d.message)
}
func div(a, b int)(int,error){
	if b  == 0 {
		return -1, &divError{
			arg: 0,
			message:"Cannot devide by 0",
		}
	}
	return a / b, nil
}

func custom_error(){
	_, err := div(10,0)
	var ae *divError

	if errors.As(err, &ae){
		fmt.Println(ae.arg)
		fmt.Println(ae.message)
	} else {
		fmt.Println("Error doe not match divError")
	}
	fmt.Println("Hello World...")
}