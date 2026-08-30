package main

func count(yield func(int) bool) {
	yield(1)
	yield(2)
	yield(3)
}

func main() {
	data := count()
}
