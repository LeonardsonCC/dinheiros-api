FROM golang:1.23-alpine AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o app-bin ./cmd/api

FROM alpine:3.20

WORKDIR /app

COPY --from=build /app/app-bin .

RUN apk --no-cache add ca-certificates tzdata

CMD ["/app/app-bin"]
