FROM golang:1.22.5 AS builder

RUN mkdir app
COPY . /app

WORKDIR /app

RUN go build -o main cmd/app/main.go

FROM golang:1.22.5

WORKDIR /app

COPY --from=builder /app .

CMD ["/app/main"]

