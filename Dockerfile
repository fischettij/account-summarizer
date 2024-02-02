FROM golang:1.21

WORKDIR /go/src/app

COPY . .

RUN go build -o summarizer ./cmd/summarizer

EXPOSE 8080

CMD ["./summarizer"]
