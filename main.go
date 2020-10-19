package main

import (
	"fmt"
	"vue_back/blog"
)

func main() {
	var err error
	blog, err := blog.MakeBlog()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer blog.CloseDB()
	blog.Run()
}
