FROM golang:1.17

ENV SRC_DIR=/go/src/github.com/kazu1029/workspace_2022/grpc/grpc-docker-todo/api

WORKDIR $SRC_DIR

COPY ./api $SRC_DIR

RUN go get -u google.golang.org/grpc \
    && go get -u github.com/golang/protobuf/protoc-gen-go \
    && go get go.mongodb.org/mongo-driver/mongo