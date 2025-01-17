FROM golang:1.23.5-alpine3.20

WORKDIR /app

COPY . /app/
RUN go build -o /app/dist/discord-notifier /app/cmd/discord-notifier/main.go
RUN go clean -modcache -cache

EXPOSE 3000

CMD ["./dist/discord-notifier"]