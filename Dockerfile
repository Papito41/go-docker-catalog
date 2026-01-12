FROM golang:1.21-alpine
WORKDIR /app
# We copy EVERYTHING at once now since 'go mod tidy' worked
COPY . .
# This builds the app
RUN go build -o main Practice.go
# This runs the app
EXPOSE 8080
CMD ["./main"]