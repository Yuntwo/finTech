# mall-go

[中文](./README-cn.md)

Always wanted to complete a fully functional open source project

Decided to develop a full-featured open source project in Go

# Use of technology
- gin、hertz
- grpc、kitex
- redis
- mysql
- mongodb
- asynq | go-queue
- amqp
- elasticsearch | gofound | zinc
- prometheus
- grafana
- jaeger
- dtm
- apisix
- wechat/alipay
- zap
- viper
- docker/docker-compose/kubernetes

# Service Functions
- [x] members
  - TODO recharge
- [x] balance
- [x] multi-merchant
  - TODO withdraw
- [ ] community⌛
- [ ] installments
- [ ] crowdfunding
- [ ] spike
- [ ] group buy
- [ ] lottery
- [ ] delivery
- [ ] fresh
- [ ] coupon
- [ ] second-hand transaction ? trade old things
- [ ] IM
- [ ] live streaming
- [ ] reward



# Project Description

The project builds scaffolding based on mixgo, which can realize flexible assembly of components. It is currently the client API interface, and admin related codes will not be implemented for the time being.

# Catalog introduction

- app:  Business code Include  api grpc mq job
- common: common components error、middleware、interceptor、tool、ctxdata
- data: runtime data
- deployments: Deploy related configuration files
- docs: Project Series Documentation
- pkg: internal package

# Gateway

The front is slb followed by apisix

# Development mode

Use the microservice development pattern. api(http) --- rpc(grpc)

rpc provides basic service implementation.

api implements service aggregation business processing.


# Log

- logstash
- filebeat

# Monitor

- prometheus

# Track

- jaeger

# pub/sub

- kafka
- mq

# Message queue、Delay queue、Timed task

- message queue
  - asynq
  - amqp
- delay queue
  - asnyq
  - amqp
- timed task
  - cron

# Distributed transaction

- dtm

# Deployment

develop use docker/docker-compose

deployment use kubernetes


# TODO

1. Add grpc checksum
2. permission check
3. grpc error handling
-[ ] 重构营销系统: 营销投放、营销活动、营销资产


# License

Apache License Version 2.0, http://www.apache.org/licenses/

Testing是随机顺序，但不是并行的，所以不会有并发问题，可以通过t.Parallel()方法设置并行测试

t.Cleanup 简化测试函数级别的清理
```
t.Cleanup(func() {
    // 测试结束时自动清理
    teardown()
})
```
推荐在每一个测试函数中单独启动测试服务器，保证测试的独立性和灵活性。虽然每次测试都启动和关闭服务器会带来一定的性能开销，但这种开销通常是可以接受的：
- httptest.NewServer启动的服务器非常轻量，开销很小。
- 如果测试函数数量较多，可以通过并行测试（t.Parallel()）来减少总运行时间。

但在某些特殊情况下可能是合理的，但尽量将全局状态的影响降到最低，并确保测试之间的隔离性：
- 测试需要依赖外部服务：例如，测试需要连接到一个真实的数据库或第三方API，而这些服务的初始化成本较高。
- 全局状态无法避免：如果测试必须依赖某些全局状态（例如全局配置或单例对象），可以在TestMain中初始化。

可以通过t.Run方法在一个测试函数中执行多个子测试，这样可以避免重复的初始化和清理工作，提高测试的效率。
但自测试之间还是独立的，只是共享一些状态，所以需要注意避免测试之间的状态污染。但本质上还是单元测试

考虑因素
- 共享资源
- 串行并行，依赖关系
- 重复的初始化和清理工作(只是从代码层面的话可以提取公共代码，但是资源还是独自创建的)


# Protobuf不仅是高效的二进制协议，还是一种接口定义语言（IDL），可以定义服务接口和消息格式，然后通过插件生成不同语言的代码，方便开发者使用。
实际rpc服务器和客户端的实现，可以使用不同的库，只要遵循protobuf定义的接口和消息格式即可。比如，可以使用gRPC、Thrift、Dubbo等框架实现rpc服务，也可以使用其他库实现rpc客户端。
还可以通过grpc-gateway生成RESTful API，提供HTTP服务器(使用的rpc代理)

- Protobuf：不仅仅是RPC接口的格式，还支持数据序列化、代码生成等功能。

# 同时实现HTTP接口和RPC接口的两种方式
1. 定义RPC服务器，HTTP服务器通过RPC客户端上下文手动调用RPC接口；适合需要灵活性和定制化的场景，但代码冗余和维护成本较高。
2. 定义RPC服务器，通过grpc-gateway自动将HTTP请求转发到RPC服务器；适合快速开发和行为一致的场景，但灵活性较低。



