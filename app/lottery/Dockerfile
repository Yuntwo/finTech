# Dockerfile语法结构：Dockerfile命令 + 子命令

# 多阶段构建：每个阶段可以用不同的基础镜像；后面阶段可以通过镜像名称引用前一个镜像构建的结果；最终只保留最后一个阶段的镜像

# 基础镜像的名称自定义
# go1.18-alpine是一个基于Alpine Linux的Go语言1.18版本的镜像，包含了Go语言的运行环境
FROM golang:1.18-alpine AS builder

# LABEL指令用于为镜像添加元数据，以键值对的形式存储，可以在镜像构建后进行查询
# 它可以用来标记或记录镜像的构建阶段、版本信息、维护者信息等，例如
# LABEL version="1.0"
# LABEL maintainer="example@example.com"
# 这里用于标记该构建阶段为gobuilder
LABEL stage=gobuilder

# 关闭了Go语言的cgo功能，确保生成的是不依赖C库的纯Go静态二进制文件，这在跨平台编译或构建 Docker 镜像时非常有用
ENV CGO_ENABLED 0
# 设置Go模块代理，用逗号分隔表示多个，按照顺序尝试
# https://goproxy.cn用于中国大陆用户，direct不使用代理而直接连接到原始Go官方模块源proxy.golang.org
ENV GOPROXY https://goproxy.cn,direct
#ENV GO111MODULE on

# RUN表示执行代码。apk是Alpine的包管理工具，update用于更新包索引，add表示安装包
# tzdata是时区数据包，no-cache表示不缓存索引和包文件
RUN apk update --no-cache && apk add --no-cache tzdata

# WORKDIR相当于mkdir -p(paraents父目录不存在的话也创建) /build创建目录 + cd /build 切换到该目录
# 可以使用相对目录，会基于当前的工作目录(通常是上一条WORKDIR设置的目录)来创建和切换到该目录
# 建议使用绝对路径来避免路径解析上的混淆，尤其是在多个WORKDIR命令串联使用时
WORKDIR /build

# COPY命令用于将本地文件复制到镜像中，COPY <源路径/相对于当前上下文目录即Dockerfile所在目录> <目标路径/容器内目录>
COPY . .

# 执行Go模块下载命令，下载所有项目的依赖包；默认开发过程中已经维持好了go.mod和go.sum文件，构建阶段不需要go mod tidy
RUN go mod download

# Shell(如Bash/sh)允许在单个命令前设置多个环境变量，且这些变量仅对该命令生效
# 格式：VAR1=value1 VAR2=value2 ... command_name，注意无需export
# 推荐这样使用，好处：
# 1、如果采用多个RUN export的方式会创建新会话，导致无法传递环境变量
# 2、仅对当前命令生效，不会污染全局环境，不会影响后续命令
# 3、避免使用ENV全局变量的方式，减少镜像层，避免变量污染
# GOOS表示目标操作系统，在交叉编译时有用。Q:不确定是否有必要指定

# go build命令
# -ldflags用于传递Go链接器的标志/参数，-s表示去除符号表信息，-w表示去除DWARF调试信息
## 符号表是二进制文件中用于存储变量、函数名等符号信息的数据结构，通常用于调试或动态链接；DWARF是一种调试信息格式，用于存储源代码和二进制代码之间的映射关系
## -s和-w都适用于生产环境减少最终输出可执行二进制文件体积；会牺牲调试能力
## -X设置程序中的变量值(常用于注入版本信息)
# -tags用于指定构建时的预设构建(限制)标签，wireinject通常用于控制启用某些特定的代码路径或功能。这里与依赖注入工具wire相关，用于生成或使用依赖注入的代码
# -installsuffix用于指定安装目录的后缀，Q:不是很懂这里这样指定有什么区别
# -o用于指定输出文件的名称，service 是生成的二进制文件的名称
# ./cmd/main.go ./cmd/wire_gen.go是构建的入口文件列表。main.go是主程序文件，wire_gen.go是由wire工具生成的依赖注入代码文件
RUN GOOS=linux CGO_ENABLED=0 go build -ldflags="-s -w" -tags wireinject -installsuffix cgo -o service ./cmd/main.go ./cmd/wire_gen.go

# 运行阶段
# Alpine是一个基于Linux的轻量级运行时环境，适合用于容器化应用
FROM alpine

# apk是Alpine的包管理工具，update用于更新包索引，add表示安装包
# --no cache表示不缓存索引和包文件；ca-certificates用于支持SSL/TLS连接，确保服务能够与外部服务器建立安全的连接(例如访问HTTPS网站)
RUN apk update --no-cache && apk add --no-cache ca-certificates

# COPY --from=builder <source> <destination>，从名为builder的镜像中复制文件到当前镜像中，构这里是将时区数据文件复制到当前镜像的相同位置，确保容器使用的是上海时区
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /usr/share/zoneinfo/Asia/Shanghai
# 文件处理还可以用ADD，用于复杂场景，支持从URL下载文件并添加到镜像中、自动解压.tar 文件等；一般推荐使用COPY
ENV TZ Asia/Shanghai

WORKDIR /app

COPY --from=builder /build/service /app/service

# 不指定--from=stage/镜像名称的话，就从build context(the directory containing the Dockerfile)复制
COPY ./etc /app/etc

# CMD表示容器启动时执行的命令，CMD ["可执行文件", "参数1", "参数2", ...]
# -f 参数指定配置文件路径
# TODO 这里写死了配置文件，多个环境下需要不同的配置文件，可以通过环境变量传递、参数传递、Docker compose等
CMD ["./service", "-f", "./etc/config.yaml"]
