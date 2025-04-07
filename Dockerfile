FROM golang:alpine AS builder

RUN apk add --no-cache ca-certificates && update-ca-certificates

WORKDIR /app
COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bot cmd/main.go

FROM scratch
COPY --from=builder /app/bot /
ENTRYPOINT ["/bot"]