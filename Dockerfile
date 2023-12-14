FROM golang:1.21.5


RUN go env -w GO111MODULE=on \
    && go env -w GOPROXY=https://goproxy.cn,direct

WORKDIR /go/cache
ADD go.mod .
ADD go.sum .
RUN go mod download

ADD . /app
WORKDIR /app


RUN CGO_ENABLED=1 go build -o main main.go && chmod +x main

CMD ["./main"]
