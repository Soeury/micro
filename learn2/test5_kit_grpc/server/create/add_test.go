package create

import (
	"context"
	"fmt"
	"micro/learn2/test5_kit_grpc/server/proto"
	"net"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

// 使用 bufconn 构建测试链接，避免使用实际端口号启动服务 这个方法可以参考一下
const bufSize = 1024 * 1024 // 1 MB 作为 bufListener 缓冲区的大小

var bufListener *bufconn.Listener // 监听内存中的连接

func init() {
	bufListener = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	// gs := NewGRPCServer(addService{})
	// proto.RegisterAddServer(s, gs)
	go func() {
		if err := s.Serve(bufListener); err != nil {
			fmt.Printf("s.serve failed:%v\n", err)
			return
		}
	}()
}

// 自定义的拨号器函数
// 使用 Dial 方法创建一个新的内存连接，这个连接将被用于gRPC客户端与服务器之间的通信
func BufDialer(context.Context, string) (net.Conn, error) {
	return bufListener.Dial()
}

func TestAdd(t *testing.T) {

	// 连接
	conn, err := grpc.NewClient(
		"bufnet",
		grpc.WithContextDialer(BufDialer),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		fmt.Printf("grpc.NewClient failed")
	}
	defer conn.Close()

	// 创建客户端
	client := proto.NewAddClient(conn)

	// 调用
	ctx, cancel := context.WithTimeout(context.Background(), time.Hour*100)
	defer cancel()

	resp, err := client.Sum(ctx, &proto.SumRequest{A: 5, B: 10})

	// 断言
	assert.Nil(t, err)
	assert.NotNil(t, resp.GetValue())
	assert.Equal(t, int64(15), resp.GetValue())
}
