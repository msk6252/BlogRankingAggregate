build:
		go mod download
		env GOOS=linux go build -ldflags="-s -w" -o bin/main ./main.go