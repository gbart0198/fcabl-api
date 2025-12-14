FROM golang:1.25.3-alpine AS builder

RUN apk add --no-cache git ca-certificates tzdata

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/api cmd/api/main.go

FROM alpine:3.23.0

RUN apk --no-cache add ca-certificates tzdata

RUN addgroup -g 1000 appuser && \
  adduser -D -u 1000 -G appuser appuser

WORKDIR /home/appuser

COPY --from=builder /app/bin/api .

RUN chown -R appuser:appuser /home/appuser

USER appuser

EXPOSE 8080

CMD ["./api"]
