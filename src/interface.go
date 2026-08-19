package main

import "fmt"

type ranjit struct {
	side int
}

// normal function
func getSide(r ranjit) string {
	fmt.Println(r)
	return "Successful"
}

func (r ranjit) name() string {
	fmt.Println("Ranjit")
	return "ranjit"
}

func main() {

	side := ranjit{side: 12}
	side.name()
	res := getSide(side)
	fmt.Println(res)

}
