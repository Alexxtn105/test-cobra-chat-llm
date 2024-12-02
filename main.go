/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/joho/godotenv"
	"test-cobra-chat-llm/cmd"
)

func main() {
	// Загружаем .env
	//godotenv.Load()
	// указываем конкретный .env-файл для загрузки
	godotenv.Load("./.env.prod")

	cmd.Execute()
}
