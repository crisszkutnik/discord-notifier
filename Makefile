start:
	go run cmd/discord-notifier/main.go

build:
	go build -o dist/discord-notifier cmd/discord-notifier/main.go

docker:
	docker build . -t ghcr.io/crisszkutnik/discord-notifier