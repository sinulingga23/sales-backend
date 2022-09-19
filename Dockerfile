FROM golang:1.19.0-alpine

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o sales_backend .

CMD ["./sales_backend"]
