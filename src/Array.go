package main
import "fmt"

func main(){
	fmt.Println("Data");
	var a [5]int;
	fmt.Println(a);
	a[1] = 100;
	fmt.Println(a);
}