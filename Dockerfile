# Этап 1: Сборка приложения
FROM golang:latest

# Установка рабочей директории
WORKDIR /app

COPY . .

EXPOSE 8080

# Команда для запуска приложения
ENTRYPOINT  ["go", "run", "."]