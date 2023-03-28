# Build Stage
FROM golang:1.20.2-alpine AS Builder
WORKDIR /app
COPY . .
RUN go build -o main main.go


# Run Stage
FROM alpine
WORKDIR /app
COPY --from=builder /app/main .
COPY app.env .

EXPOSE 8080
CMD  [ "/app/main" ]