PROTO_OUTPUT_FILES = rpc/multiplier/service.pb.go rpc/multiplier/service.twirp.go

.PHONY: proto
proto: ${PROTO_OUTPUT_FILES}

${PROTO_OUTPUT_FILES}: rpc/multiplier/service.proto
	protoc --proto_path=${GOPATH}/src:. --twirp_out=. --go_out=. $^


# Note: official docs recommend `retool`: https://twitchtv.github.io/twirp/docs/install.html#get-protoc-gen-go-and-protoc-gen-twirp-plugins
setup:
	go get github.com/golang/protobuf/protoc-gen-go
	go get github.com/twitchtv/twirp/protoc-gen-twirp

.PHONY: server
server:
	go run cmd/server/*.go

.PHONY: client
client:
	go run cmd/client/*.go
