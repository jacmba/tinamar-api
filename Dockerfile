FROM golang:alpine

RUN mkdir -p /src
RUN mkdir -p /app

WORKDIR /src

COPY . /src

RUN apk add --no-cache git

RUN go get -v
RUN go build -o /app/tinamar-api
WORKDIR /app

CMD ./tinamar-api