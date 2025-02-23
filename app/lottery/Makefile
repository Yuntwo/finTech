GOHOSTOS:=$(shell go env GOHOSTOS)
GOPATH:=$(shell go env GOPATH)
#  Git命令git describe --tags --always获取当前项目的版本号,这个版本号通常是基于Git标签生成的
VERSION=$(shell git describe --tags --always)

# .PHONY(phony target)声明伪目标，确保无论文件/目标是否存在，make都会执行相应的命令
# 伪目标通常用于执行某些命令而不是生成文件，因为Makefile默认以编译生成文件为目标，所以不生成文件的话需要声明伪目标
# Makefile基本语法是：
# target(文件/对象名): dependencies(依赖文件/对象名)
#    command
# 解析规则是：如果存在且目标文件比依赖文件都新，make就认为目标已经是最新的，不会执行命令；如果目标不存在，或者目标依赖的文件比目标新，make会执行命令来更新目标
# 这种伪目标的规则有点冗余
.PHONY: init
# init env
init:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/google/gnostic/cmd/protoc-gen-openapi@latest
	go install github.com/google/wire/cmd/wire@latest


.PHONY: run
run:
	go run ./...

.PHONY: build
build:
	# TODO 弄懂wire依赖注入工具
	CGO_ENABLED=0 go build -ldflags "-X main.Version=$(VERSION)" -tags wireinject -o service ./cmd/main.go ./cmd/wire_gen.go

.PHONY: api
api:
	buf generate

.PHONY: generate
# generate
generate:
	go get github.com/google/wire/cmd/wire@latest
	go generate ./...
