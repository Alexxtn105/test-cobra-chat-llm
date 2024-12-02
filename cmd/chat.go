/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

// chatCmd represents the chat command
var chatCmd = &cobra.Command{
	Use:   "chat",
	Short: "LLM чат-бот",
	Long:  `LLM чат-бот`,
	Run: func(cmd *cobra.Command, args []string) {
		// здесь необходимо вернуть введенные пользователем данные
		//fmt.Println("chat called")
		reader := bufio.NewReader(os.Stdin)

		// логика "Мягкого завершения
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

		go func() {
			<-sigChan
			fmt.Println("Перехвачен сигнал прерывания исполнения. Выходим...")
			os.Exit(0)
		}()

		//создаем бесконечный цикл, в котором пользователь может вводить текст
		for {
			// пишем строку приглашения
			fmt.Print("> ")
			// читаем введенный пользователем текст (до символа "Ввод" - \n)
			input, _ := reader.ReadString('\n')
			// триммируем строку
			input = strings.TrimSpace(input)

			switch strings.ToLower(input) {
			case "quit", "exit":
				//выходим из программы
				fmt.Println("Выход из программы...")
				os.Exit(0)
			default:
				fmt.Println("Вы сказали: ", input) // отправляем пользователю введенную строку
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(chatCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// chatCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// chatCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
