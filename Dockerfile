FROM golang:latest

LABEL maintainer="Orwyn Carvalho <orwyn.carvalho@gmail.com>"

WORKDIR /app

# COPY go.sum .

# COPY go.mod .

# RUN go mod download

COPY . .

RUN go get -u github.com/golang/protobuf

RUN go get -u google.golang.org/grpc

ENV PORT 50051

RUN go build user/user_server/server.go user/user_server/helper.go user/user_server/db.go user/user_server/models.go
# RUN GOPATH=$GOPATH:/root go build user/user_server/*.go


CMD ["./server"]