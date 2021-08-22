GO = GO111MODULE=on GOPROXY="https://goproxy.io,direct" GOSUMDB="sum.golang.google.cn" go
PBDIR = ./rpc
SRCS_PB += $(foreach dir, $(PBDIR), $(wildcard $(dir)/*.pb.go))

rpc:
	@protoc -I ./api ./api/*.proto --go_out=plugins=grpc:${PBDIR}
	@for pb in ${SRCS_PB}; do protoc-go-inject-tag -input=$$pb; done

build:rpc
	# TODO

image:rpc
	# TODO

deploy:rpc
	# TODO

.PHONY: rpc build image deploy
