FROM golang:alpine

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

RUN go version
ENV GOPATH=/

COPY .. ./

RUN go mod download
RUN go build -o jinder-api ./cmd/main.go

RUN chmod +x ./jinder-api
CMD ["./jinder-api"]
