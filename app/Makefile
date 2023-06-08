build:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o main *.go

run:
	go run main.go

zip: build
	zip main.zip main
	rm main

deploy: zip
	aws lambda update-function-code --function-name diete-bot --zip-file fileb://main.zip
	rm main.zip
	@echo "deployed!"