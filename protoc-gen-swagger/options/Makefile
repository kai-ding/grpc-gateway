all:
	protoc -I/usr/local/include -I. \
		-I${GOPATH}/src \
		-I${GOPATH}/src/github.com/kai-ding/grpc-gateway \
		-I${GOPATH}/src/github.com/kai-ding/grpc-gateway/third_party/googleapis \
		--go_out=. \
		annotations.proto