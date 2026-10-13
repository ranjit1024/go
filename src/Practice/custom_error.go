package main
import (
	"fmt"
	"errors"
)

type divError struct{
	num int
	message string
}

func (d *divError) Error() string{
	return fmt.Sprintf("%d - %s", d.num, d.message)
}

func div( a, b int) (int, error){
	if b == 0 {
		return -1, &divError{
			num: 0,
			message:"Can not devide with 0",
		}
	}
	return a / b, nil
}

func custom_error(){
	_ , err := div(12,0)
	var ae *divError
	if errors.As(err, &ae){
		fmt.Println(ae.num)
		fmt.Println(ae.message)
	}
	fmt.Println("Hello World...")
}