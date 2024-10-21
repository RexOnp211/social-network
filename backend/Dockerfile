FROM golang:1.22-bookworm

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .

WORKDIR /app/cmd/api

RUN go build -o main .

CMD ["/app/cmd/api/main"]
