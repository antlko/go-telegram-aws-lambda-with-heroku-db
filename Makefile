build:
	export GO111MODULE=on
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/app app/main.go

clean:
	rm -rf bin

deploy: clean build
	sls deploy