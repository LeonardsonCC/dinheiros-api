FROM golang:1.22-alpine3.19

WORKDIR /src

COPY . .

RUN go mod download

RUN go build -o dinheiros ./cmd/api/

CMD [ "/src/dinheiros" ]
