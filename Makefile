build:
	go build -o bin/pricefetcher.exe

run: build
	.\bin\pricefetcher.exe

