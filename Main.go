package main

import (
	"fmt"

	"tudai2021.com/model"
)

func main() {
	Result, err := model.Parser("TX05EPQKR")
	fmt.Println(Result, err)
}
