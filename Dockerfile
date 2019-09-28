FROM golang:1.13.1-alpine

WORKDIR /app/src
COPY ./backend/src .

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

RUN go get -d -v ./...
RUN go build -o ../build/one-time-pad-server

CMD ["../build/one-time-pad-server"]