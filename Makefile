build:
	go build -o server cmd/app/main.go

run: build
	go run cmd/app/main.go

watch:
	reflex -s -r '\.go$$' make run

