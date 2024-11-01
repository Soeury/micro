package consul

import (
	"fmt"
	"net"

	"github.com/hashicorp/consul/api"
)

type Consul struct {
	Client *api.Client
}

// 连接 consul
func ConnectToConsul() (*Consul, error) {

	cc, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		return nil, err
	}
	return &Consul{Client: cc}, nil
}

// 注册服务
func (c *Consul) RegisterService(serName string, ip string, port int) error {

	// 定义健康检查
	check := &api.AgentServiceCheck{
		GRPC:                           fmt.Sprintf("%s:%d", ip, port),
		Interval:                       "10s",
		Timeout:                        "10s",
		DeregisterCriticalServiceAfter: "60s",
	}

	// 定义服务
	srv := &api.AgentServiceRegistration{
		ID:      fmt.Sprintf("%s-%s-%d", serName, ip, port),
		Name:    serName,
		Tags:    []string{"rabbit", "multi"},
		Port:    port,
		Address: ip,
		Check:   check,
	}

	return c.Client.Agent().ServiceRegister(srv)
}

// DeregisterService 退出时撤销服务
func (c *Consul) DeregisterService(serviceID string) error {
	return c.Client.Agent().ServiceDeregister(serviceID)
}

// GetLocalAddr 获取本机ip
func GetLocalAddr() (net.IP, error) {

	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	addr := conn.LocalAddr().(*net.UDPAddr)
	return addr.IP, nil
}
