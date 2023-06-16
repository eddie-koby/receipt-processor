FROM golang:1.20

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY *.go ./
COPY models/ ./models/
COPY examples/ ./examples/

RUN go build -o /receipt-processor

EXPOSE 8080

CMD [ "/receipt-processor"]