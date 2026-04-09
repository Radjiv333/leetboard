run: 
	go run cmd/main.go

drun: 
	docker build -t leetboard .
	docker run -p 8080:8080 leetboard