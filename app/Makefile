build:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o main main.go

run:
	go run main.go

zip: build
	zip main.zip main
	rm main