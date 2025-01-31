build:
	cd client && pnpm build
	ENV=prod go build -o bin/encurtago cmd/api/main.go

dev:
	cd client && pnpm dev & air && fg
