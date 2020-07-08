FROM golang:1.14

WORKDIR /code/dockling
COPY . .

RUN go build -o bin/dockling ./cmd/dockling

CMD ["bin/dockling"]