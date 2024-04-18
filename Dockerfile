FROM golang:1.21.6

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . .

RUN go build -o main .

ENV PORT=8080

EXPOSE 8080

CMD ["./main"]
