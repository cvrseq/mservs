FROM golang:1.23.1

WORKDIR /microservices

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o /go-app ./api-service/cmd

EXPOSE 8080

CMD ["/go-app"]