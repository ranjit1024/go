package main

import "fmt"

type oparaionError struct {
	oparator int
	message  string
}

func (SumError *oparaionError) Error() string {
	return fmt.Sprintf("%d - %s", SumError.oparator, SumError.message)
}

func add(a, b int) (int, error) {
	if a == 0 || b == 0 {
		if a == 0 {
			return -1, &oparaionError{oparator: a, message: "zero is not allowed"}
		}
		return -1, &oparaionError{oparator: b, message: "Zero is not allowed"}
	}
	return a + b, nil
}

func main() {
	fmt.Println("Data is the oil")
	result, err := add(12, 12)
	fmt.Println(result)
	fmt.Println(err)
}
