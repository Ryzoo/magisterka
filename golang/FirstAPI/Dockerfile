FROM golang:alpine

RUN mkdir /app
COPY . /app

WORKDIR /app
RUN go build -o startup main.go

CMD ["/app/main"]