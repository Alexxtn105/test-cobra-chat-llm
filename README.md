# Тестовый проект чат-бота с с использованием Cobra и LangChain

## Установка необходимых зависимостей
Установка Cobra:
```bash
go get -u github.com/spf13/cobra@latest
```
Установка Cobra-CLI (для генерации команд):
```bash
go install github.com/spf13/cobra-cli@latest
```
Установка LangChainGo:
```bash
go get github.com/tmc/langchaingo
```

## Инициализация Cobra.
В результате будет создана структура программы - файлы `main.go` и `./cmd/root.go` со сгенерированным кодом.
```bash
cobra-cli init
```

Добавить новую команду, например, `chat`:
```bash
cobra-cli add chat
```
В результате будет сгенерирован код команды в файле `./cmd/chat.go`

Запуск чата:
```bash
go run main.go chat
```

Выход из чата - команда `exit` или `quit`