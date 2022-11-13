package main

import (
	"fmt"
	"os"
)

func initSetEnv() {
	os.Setenv("PATH", os.Getenv("PATH"))
}

func main() {
	fmt.Println("encode server")

	// 初期化
	initSetEnv()
}

func encode(inputFileName string) {
	fmt.Println(inputFileName)
}
