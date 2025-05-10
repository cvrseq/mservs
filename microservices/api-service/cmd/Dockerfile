FROM golang:1.24.2

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o /go-app ./api-service/cmd

EXPOSE 8000

CMD ["/go-app"]