package main
import (
	"fmt"
	"iter"
)

func coutToFive() iter.Seq[int]{
	return func(yield func(int) bool){
		for i := 1; i <= 5; i++{
			if !yield(i){
				return
			}
		}
	}
}

func demo_range_over_itarator(){
	for n := range coutToFive(){
		fmt.Println(n)
	}
}