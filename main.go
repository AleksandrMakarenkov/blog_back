package main

import (
	"fmt"
	"vue_back/blog/dependencies"
)

func main() {
	var err error
	blog, err := dependencies.MakeBlog()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer blog.CloseDB()
	blog.Run()
}
