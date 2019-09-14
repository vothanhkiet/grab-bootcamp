stub:
	protoc \
		-I/usr/local/include \
		-I. \
		-I${GOPATH}/src \
		-I${GOPATH}/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.9.5/third_party/googleapis \
		-I${GOPATH}/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.9.5 \
		--go_out=plugins=grpc:./ \
		./transport/service.proto

proxy:
	protoc \
		-I/usr/local/include \
		-I. \
		-I${GOPATH}/src \
		-I${GOPATH}/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.9.5/third_party/googleapis \
		-I${GOPATH}/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.9.5 \
		--grpc-gateway_out=logtostderr=true:./ \
		./transport/service.proto
swagger:
	protoc \
		-I/usr/local/include \
		-I. \
		-I${GOPATH}/src \
		-I${GOPATH}/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.9.5/third_party/googleapis \
		-I${GOPATH}/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.9.5 \
		--swagger_out=logtostderr=true:./ \
		./transport/service.proto
	mv ./transport/service.swagger.json docs/swagger.json

omitempty:
	ls ./transport/*.pb.go | xargs -n1 -IX bash -c "sed -e '/bool/ s/,omitempty//' X > X.tmp && mv X{.tmp,}"
	ls ./transport/*.pb.go | xargs -n1 -IX bash -c "sed -e '/int32/ s/,omitempty//' X > X.tmp && mv X{.tmp,}"
	ls ./transport/*.pb.go | xargs -n1 -IX bash -c "sed -e '/\[]/ s/,omitempty//' X > X.tmp && mv X{.tmp,}"
	ls ./transport/*.pb.go | xargs -n1 -IX bash -c "sed -e '/*Paging / s/,omitempty//' X > X.tmp && mv X{.tmp,}"
	
all: stub proxy swagger