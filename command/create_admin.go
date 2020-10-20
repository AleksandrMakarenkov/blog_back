package main

import (
	"bufio"
	"fmt"
	_ "github.com/jackc/pgx/stdlib"
	"os"
	"strings"
	"vue_back/blog/dependencies"
	"vue_back/blog/model"
	hasher "vue_back/blog/service/password"
)

func main() {
	fmt.Println("Please input email (should be unique):")
	reader := bufio.NewReader(os.Stdin)
	email, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}
	email = strings.TrimSuffix(email, "\n")

	fmt.Println("Please enter password:")
	password, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}
	password = strings.TrimSuffix(password, "\n")

	reformDB, err := dependencies.MakeReform()
	if err != nil {
		fmt.Println(err)
		return
	}

	password, err = hasher.Hash(password)
	if err != nil {
		fmt.Println(err)
		return
	}

	user := model.User{
		Email:    email,
		Password: password,
		Role:     "admin",
	}
	err = reformDB.Insert(&user)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Admin created!")
}
