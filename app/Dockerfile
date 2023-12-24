FROM golang:1.21.5-alpine
# RUN apk add build-base

WORKDIR /app
COPY ./go.mod ./go.sum ./
RUN go mod download

COPY . .

RUN go build -o app .
EXPOSE 8080

CMD ["/app/app"]
