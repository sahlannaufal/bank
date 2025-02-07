FROM golang:1.21-alpine
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o main .
CMD ["sh", "-c", "./main migrate && ./main"]
