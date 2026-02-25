FROM golang:1.25.7-alpine AS builder

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . .
RUN go build -o myapp ./cmd/api/main.go

FROM alpine
COPY --from=builder app/database/migrations/ database/migrations
COPY --from=builder app/myapp /myapp
ENTRYPOINT ["/myapp"]