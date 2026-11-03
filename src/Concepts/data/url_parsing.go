package main

import (
	"fmt"
	"net/url"
)

var f = fmt.Println

func url_parsing() {
	s := "postgres://user:pass@host.com:5432/path?k=v#f"
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}
	f(u.User.Password())
	f(u.Path)
	f("Hello World")
}
