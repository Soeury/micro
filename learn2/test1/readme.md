## 服务注册与服务发现


# 实现方案
1. 服务注册
    真正提供服务的实例的模块信息注册到一个公共中心上面

2. 服务发现
    -1. 服务及其调用方直接与注册中心交互
        ---- 优势：支持多平台，格式类型的服务，不依赖于某个平台的部署
        ---- 缺点：需要在客户端和服务器端集成 SDK，并且需要维护相关代码逻辑，需考虑到 SDK 对于不同语言的支持情况

    -2. 通过部署基础设施来完成服务发现    
        ---- 优势: 所有服务发现内容都由平台处理，客户端和服务器端不包含任何服务发现的代码
        ---- 缺点: 仅限于使用该平台的部署服务，多依次网络跳转，存在性能损耗


# 主流注册中心
1. Eureka
2. Zookeeper
3. Consul
4. Etcd
...


# CAP理论
1. consistency ：一致性，指的是在分布式系统中，数据在多个副本之间能够保持一致的状态
2. availability : 可用性，指在分布式系统中，对于每一个请求在有限的时间内都会给出响应，而不是无期限的等待
3. partition tolerance : 分区容忍性，指分布式系统在遇到网络分区时，仍然能够保持系统的正常运行。

**** 网络分区： 指分布式系统中的多个节点之间由于网阔故障无法进行通信
**** 一致性 ：数据在还未同步的情况下阻塞请求，等到数据完全同步才返回响应，保证了数据的一致性
**** 可用性 ：数据在未同步的情况下也会返回响应，只不过返回的数据可能是旧数据，保证了数据的高可用性
**** 分区容忍性 ：在分布式系统中，分区容忍性是必备的特性


# Raft协议
