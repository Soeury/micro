package consul

import (
	"fmt"

	"github.com/hashicorp/consul/api"
)

// 通过 api 查询 consul 来完成服务发现，但是我们一般不会这么去做
// 使用第三方的库 _ "github.com/mbobakov/grpc-consul-resolver" 来完成服务发现
// 按照上面的第三方库的方式就不需要写下面这些代码了，只需要改动 addr 即可，注意匿名导入包
func GetConsulAddr() (string, error) {

	// 1. 连接 consul
	cc, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		fmt.Printf("newClient failed:%v\n", err)
		return "", err
	}

	// 2. 寻找服务
	serviceMap, err := cc.Agent().ServicesWithFilter("learn2-test2")
	if err != nil {
		fmt.Printf("cc.agent.serviceWithFilter failed:%v\n", err)
		return "", err
	}

	// 3. 从返回的服务中选出需要的
	var addr string
	for k, v := range serviceMap {
		fmt.Printf("%s-%+v\n", k, v)
		addr = fmt.Sprintf("%s:%d\n", v.Address, v.Port)
		break
	}

	return addr, nil
}
