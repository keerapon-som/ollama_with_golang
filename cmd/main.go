package main

import (
	"fmt"
	"ollamawithgo/config"
	"ollamawithgo/internal/handler"
)

func main() {
	config := config.GetConfig()
	fmt.Println(config.Ollama.BASEURL)
	fmt.Println("Hello, World!")

	app := handler.CreateHandler()
	err := app.Listen(":8080")
	if err != nil {
		panic(err)
	}

}
