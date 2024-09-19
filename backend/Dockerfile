FROM golang:1.22 AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download
COPY . ./

WORKDIR /app/cmd/api
RUN go build -o /bin/myapp .

FROM debian:bookworm AS run
WORKDIR /app
COPY --from=builder /bin/myapp /bin/myapp
EXPOSE 8080

CMD ["/bin/myapp"]
