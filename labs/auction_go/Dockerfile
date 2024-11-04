FROM golang:1.22

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o /app/auction cmd/auction/main.go

EXPOSE 8080

ENTRYPOINT ["/app/auction"]