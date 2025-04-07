# Используем официальный образ Go для сборки
FROM golang:1.19 as builder

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Копируем go.mod и go.sum для загрузки зависимостей
COPY go.mod go.sum ./

# Загружаем зависимости
RUN go mod download

# Копируем остальные файлы проекта
COPY . .

# Собираем бинарный файл
RUN go build -o main cmd/main.go

# Используем минимальный образ для запуска
FROM alpine:latest

# Устанавливаем рабочую директорию
WORKDIR /root/

# Копируем бинарный файл из предыдущего этапа
COPY --from=builder /app/main .

# Запускаем приложение
CMD ["./main"]
