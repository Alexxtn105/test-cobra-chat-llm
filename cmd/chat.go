/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/openai"
	"log"
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

		// подключаем LLM
		llm, err := openai.New()
		if err != nil {
			log.Fatal(err)
		}

		// берем контекст
		ctx := context.Background()

		// Вводим инициализирующую строку
		fmt.Print("Введите инициализирующую строку для LLM: ")
		// читаем введенное пользователем
		initialPrompt, _ := reader.ReadString('\n')
		initialPrompt = strings.TrimSpace(initialPrompt)
		// инициализируем слайс для хранения сообщений
		content := []llms.MessageContent{
			llms.TextParts(llms.ChatMessageTypeSystem, initialPrompt),
		}
		//fmt.Println(content)
		// Уведомляем пользователя, что инициализирующая строка принята
		fmt.Println("Инициализирующая строка принята. Переходим в чат...")

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
				// Process user input with the LLM here
				response := ""
				content = append(content, llms.TextParts(llms.ChatMessageTypeHuman, input))
				_, err := llm.GenerateContent(ctx, content,
					llms.WithMaxTokens(1024),
					llms.WithStreamingFunc(func(ctx context.Context, chunk []byte) error {
						fmt.Print(string(chunk))
						response = response + string(chunk)
						return nil
					}),
				)
				//fmt.Println(resp)
				if err != nil {
					fmt.Println("error:" + err.Error())
				} else {
					content = append(content, llms.TextParts(llms.ChatMessageTypeSystem, response))
				}
				//// обрабатываем введенные пользователем данные с помощью LLM
				//response := ""
				////fmt.Println("Вы сказали: ", input) // отправляем пользователю введенную строку
				//
				//// дописываем пользовательское сообщение как от человека (human)
				//content = append(content, llms.TextParts(llms.ChatMessageTypeHuman, input))
				//// генерим контент
				//llm.GenerateContent(
				//	ctx,
				//	content,
				//	llms.WithMaxTokens(1024),
				//	llms.WithStreamingFunc(func(ctx context.Context, chunk []byte) error {
				//		fmt.Print(string(chunk))
				//		response += string(chunk)
				//		return nil
				//	}),
				//)
				//
				//// дописываем сообщение как системное (system)
				//content = append(content, llms.TextParts(llms.ChatMessageTypeSystem, response))

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
