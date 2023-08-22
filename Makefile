dev:
	go run main.go

dev-release:
	env GIN_MODE=release go run main.go

build:
	env GOARCH=amd64 GOOS=linux GIN_MODE=release go build -ldflags="-s -w" -o bin/main main.go

clean:
	rm -rf ./bin

clean-build: clean build

docker-build:
	docker build -t eleutheria/project-a-api:latest .
	
docker-push:
	docker push eleutheria/project-a-api:latest