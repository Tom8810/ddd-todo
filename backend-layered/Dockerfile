FROM golang:1.24-alpine

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

# Install air for hot reloading (compatible with Go 1.24)
RUN go install github.com/air-verse/air@v1.61.0

EXPOSE 8080

CMD ["air", "-c", ".air.toml"]