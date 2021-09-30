FROM golang:1.16-alpine as builder

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o /server

FROM alpine

ENV ENVIRONMENT=development
ENV PORT=3001
ENV MYSQL_STRING=root:root@tcp(172.17.0.2:3306)/api-notificacao?charset=utf8&parseTime=True&loc=Local

WORKDIR /app

COPY --from=builder /server /server

CMD ["/server"]