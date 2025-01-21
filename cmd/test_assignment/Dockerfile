# build stage
FROM golang:1.23.2-alpine AS build

# Установите Git
RUN apk add --no-cache git

WORKDIR /src
COPY go.mod go.sum ./
# Если вы используете vendor, закомментируйте следующую строку
# RUN go mod download
COPY vendor vendor ./
COPY . .

# Сборка приложения
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-X 'main.Version=$(git rev-parse --short HEAD)'" -o /bin/test_assignment ./cmd/test_assignment

# prod stage
FROM alpine

# Копируем собранное приложение
COPY --from=build /bin/test_assignment /bin/test_assignment
COPY --from=build /src/internal/database/postgresql/migrations /migrations

# Укажите правильный путь к исполняемому файлу
ENTRYPOINT ["/bin/test_assignment"]