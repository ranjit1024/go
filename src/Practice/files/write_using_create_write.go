package main

import (
	"log"
	"os"
)

func write_using_crate_and_write() {

	// data := []byte(message)
	// err := os.WriteFile("./tmp/a.txt", data, 0644)
	// check(err)
	// fmt.Println(message)
	// fmt.Println(message)
	// file, err := os.Create("./tmp/data.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// defer file.Close()
	file, err := os.OpenFile(
		"./tmp/a.txt",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0644,
	)
	check(err)
	defer file.Close()
	_, err = file.Write([]byte("New content"))
	if err != nil {
		log.Fatal(err)
	}
	check(err)
}
