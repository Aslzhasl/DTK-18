FROM golang:1.23.4

WORKDIR /app

COPY go.mod ./

COPY go.sum ./

RUN go mod download

COPY . .


RUN go build -o app ./cmd

EXPOSE 8083

CMD [ "./app"]
