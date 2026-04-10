FROM golang:1.25.6

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . .

RUN go build -o leetboard cmd/main.go

EXPOSE 8080

ENTRYPOINT [ "./leetboard" ]

CMD ["--port", "8080"]