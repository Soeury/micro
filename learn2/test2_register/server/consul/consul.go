package consul

import (
	"fmt"
	"net"

	"github.com/hashicorp/consul/api"
)

// 步骤：
//    -1. 连接到 consul
//    -2. 定义并注册服务到 consul
//    -3. 添加健康检查
//          ---- 注册健康检查
//          ---- 配置健康检查策略
//          ---- 更改监听ip : localhost:8080  -> 0.0.0.0:8080
//    -4. client端通过第三方的库实现服务发现
//    -5. 手动服务撤销

type Consul struct {
	Client *api.Client
}

// NewConsul 连接至 consul 并返回一个 consul 对象
func NewConsul() (*Consul, error) {

	cc, err := api.NewClient(api.DefaultConfig()) // default: 127.0.0.1:8500
	if err != nil {
		return nil, err
	}
	return &Consul{Client: cc}, nil
}

// RegisterService 定义并注册服务到 consul + 健康检查
func (c *Consul) RegisterService(serviceName string, ip string, port int) error {

	// 定义健康检查策略，告诉consul如何进行健康检查
	// 定义好之后注册到下面的服务中
	check := &api.AgentServiceCheck{
		GRPC:                           fmt.Sprintf("%s:%d", ip, port), // 这里必须是外部可以访问的ip地址
		Timeout:                        "10s",                          // 超时时间
		Interval:                       "10s",                          // 检查频率
		DeregisterCriticalServiceAfter: "60s",                          // 指定时间后自动注销不健康的服务节点
	}

	// 定义服务，注意类型
	srv := &api.AgentServiceRegistration{
		ID:      fmt.Sprintf("%s-%s-%d", serviceName, ip, port), // 服务唯一ID
		Name:    serviceName,                                    // 服务名称
		Tags:    []string{"rabbit", "hello"},                    // 为服务打标签
		Address: ip,
		Port:    port,
		Check:   check, // 使用自定义的健康策略
	}

	// 注册服务，返回error
	return c.Client.Agent().ServiceRegister(srv)
}

// Deregister 退出时撤销服务
func (c *Consul) Deregister(serviceID string) error {
	return c.Client.Agent().ServiceDeregister(serviceID)
}

// GetOutIP 获取本地地址
func GetOutIP() (net.IP, error) {

	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP, nil
}
