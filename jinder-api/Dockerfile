FROM golang:1.14-buster

RUN go version
ENV GOPATH=/

COPY ./ ./

# install psql
RUN apt-get update
RUN apt-get -y install postgresql-client

# build go app
RUN go mod download
RUN go build -o jinder-api ./cmd/main.go

RUN chmod +x ./jinder-api

CMD ["./jinder-api"]
