FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY main.go .
RUN go build main.go

FROM scratch
COPY --from=builder /app/main .
CMD ["./main"]

# docker build -t go-fundamentals:0 .
# docker images | grep go-fundamentals
# docker run -p 8080:8080 go-fundamentals:0
