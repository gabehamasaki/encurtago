.PHONY: build dev run-test

build:
	@cd client && pnpm build
	@ENV=prod go build -o bin/encurtago cmd/api/main.go

dev:
	@npx concurrently -k -c "#93c5fd,#c4b5fd" "cd client && pnpm dev" "air" --names=vite,api

run-test:
	@go test ./test -cover -v
