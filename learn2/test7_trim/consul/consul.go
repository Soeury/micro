package consul

import (
	"fmt"
	"net"

	"github.com/hashicorp/consul/api"
)

type Consul struct {
	Client *api.Client
}

// 连接到 consul 并返回一个 consul 对象
func NewConsul() (*Consul, error) {

	cc, err := api.NewClient(api.DefaultConfig()) // default : 127.0.0.1:8500
	if err != nil {
		return nil, err
	}
	return &Consul{Client: cc}, nil
}

// 定义服务 注册到 Consul + 健康检查
func (c *Consul) RegisterService(serviceName string, ip string, port int) error {

	// 定义健康检查
	check := &api.AgentServiceCheck{
		GRPC:                           fmt.Sprintf("%s:%d", ip, port),
		Timeout:                        "10s",
		Interval:                       "10s",
		DeregisterCriticalServiceAfter: "30s",
	}

	// 定义服务
	srv := &api.AgentServiceRegistration{
		ID:      fmt.Sprintf("%s-%s-%d", serviceName, ip, port),
		Name:    serviceName,
		Tags:    []string{"rabbit", "go-kit"},
		Address: ip,
		Port:    port,
		Check:   check,
	}

	// 注册服务，返回 error
	return c.Client.Agent().ServiceRegister(srv)
}

// Deregister 退出时撤销服务
func (c *Consul) Deregister(serviceID string) error {
	return c.Client.Agent().ServiceDeregister(serviceID)
}

// 获取本地地址
func GetLocalIP() (net.IP, error) {

	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP, nil
}
