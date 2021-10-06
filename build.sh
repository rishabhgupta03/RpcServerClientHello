export GOPATH=/Users/rishabh/Desktop/Golang
export PATH=$PATH:$GOPATH/bin
docker run --rm -v $(pwd):$(pwd) -w $(pwd) protoc-gen:latest --go_out=plugins=grpc,paths=source_relative:./client -I. ./proto/*.proto

