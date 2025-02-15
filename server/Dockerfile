FROM golang:1.23.4 AS builder

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o /bin/server ./cmd/main.go

FROM alpine:latest

WORKDIR /src

COPY --from=builder /bin/server /bin/server

EXPOSE 5000

RUN chmod +x /bin/server

ENTRYPOINT ["/bin/server"]
