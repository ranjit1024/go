package main
import (
	"fmt"
	"errors"
)

type divError struct{
	num int
	message string
}

func (d *divError) String() string{
	fmt.Sprintf("%d - %s", d.num, d.message)
}

func div( a, b int) (int, error){
	if b == 0 {
		return -1, &divError{
			num:0,
			message:"Can not devide with 0"
		}
	}
	return a / b, nil
}
func main(){
	fmt.Println("Hello World...")
}