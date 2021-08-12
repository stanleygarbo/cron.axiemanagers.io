FROM golang:1.16.2

LABEL maintainer="Stanley Garbo<stanleygarbo@gmail.com>"

WORKDIR /app

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

RUN go build

CMD ["./08_cron.axiemanagers.io"]