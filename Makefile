run: 
	go run cmd/main.go

drun: 
	docker build -t leetboard .
	docker run -p 8080:8080 leetboard --port 8080

build:
	go build -o leetboard cmd/main.go

dbuild:
	docker build -t leetboard .

up: 
	docker compose up --build

nuke: 
	docker compose down -v 