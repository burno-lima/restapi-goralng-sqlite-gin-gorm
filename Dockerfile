FROM golang:1.19

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download
 
COPY . .

RUN CGO_ENABLED=1 GOOS=linux go build -o /api-books ./main.go

EXPOSE 8080

CMD ["/api-books"]
