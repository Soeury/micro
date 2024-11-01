package service

import (
	"context"
)

// service 层: 放置所有与业务逻辑相关的内容
//  -1. 定义接口
//  -2. 定义结构体
//  -3. 结构体实现接口
//  -4. 接口的而构造函数(初始化结构体，返回接口，用于在main中生成服务)

type ComputerService interface {
	Add(context.Context, int64, int64) (int64, error)
	Append(context.Context, string, string) (string, error)
}

// 这里可以按需添加一些自己的字段
// 比如 数据库初始化的 DB ...
type computerService struct {
	// db db.Conn
}

// 下面就是编写业务逻辑
func (computerService) Add(_ context.Context, num1 int64, num2 int64) (int64, error) {

	// 1. 查询数据
	// 2. 处理数据
	// 3. 返回数据
	return num1 + num2, nil
}

// 假设这个 Append 需要调用另外的服务来实现 ：去掉字符串内部的所有空格的服务
func (computerService) Append(_ context.Context, str1 string, str2 string) (string, error) {

	// 1. 查询数据
	// 2. 处理数据
	// 3. 返回数据
	return str1 + str2, nil
}

// NewService : addService 的构造函数
func NewService() ComputerService {

	return &computerService{
		// db...
	}
}
