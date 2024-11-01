package main

import (
	"fmt"
	"micro/learn1/test7/book"

	"github.com/iancoleman/strcase"
	fieldmask_utils "github.com/mennanov/fieldmask-utils"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

// wrapper 使用示例 ： 解决客户端传[零值]或者[不传值]的问题
// field_mask 使用示例 ： 更新时自动记录更新的字段

// 使用 fieldMaskDemo 实现部分更新字段而不是全部更新的示例
func fieldMaskDemo() {

	// client
	/*
		od := &book.Book{
			Title:     "beauitful love",
			Author:    "kaka",
			Price:     &wrapperspb.Int64Value{Value: 6699},
			SalePrice: &wrapperspb.DoubleValue{Value: 7733},
			Content:   &wrapperspb.StringValue{Value: "love you"},
			Another: &book.Book_Info{
				Teacher: "Candy",
				Art:     "Rabbit",
			},
		}
	*/

	paths := []string{"Price", "Another.teacher"} // 更新的字段信息需要提前记录下来，用[]string记录路径
	nd := &book.UpdataBook{
		Op: "admin",
		Book: &book.Book{
			Price: &wrapperspb.Int64Value{Value: 8855},
			Another: &book.Book_Info{
				Teacher: "boliang",
			},
		},
		UpdateMask: &fieldmaskpb.FieldMask{Paths: paths},
	}

	// server
	mask, _ := fieldmask_utils.MaskFromProtoFieldMask(nd.UpdateMask, strcase.ToCamel)
	var bookDst = make(map[string]interface{})

	// 将数据读取到map[string]interface{}
	// fieldmask-utils支持读取到结构体等
	// 将需要映射的数据存储成指定数据形式的 kv
	fieldmask_utils.StructToMap(mask, nd.Book, bookDst)
	for idx, v := range bookDst {
		// 这里拿到数据后就可以更新数据库
		fmt.Println(idx, v)
	}
}

func wrapDemo() {

	// client 传值
	book := &book.Book{
		Title:     "beauitful",
		Author:    "boliang",
		Price:     &wrapperspb.Int64Value{Value: 6699},
		SalePrice: &wrapperspb.DoubleValue{Value: 7733},
		Content:   &wrapperspb.StringValue{Value: "boliang is beautiful"},
	}

	// server 取值
	// 可以使用 nil 判断用户是否传值(应该也是指针类型的)
	// 平时判断用户是否为结构体字段传值时，可以使用指针类型的数据:  *int64  *string   *bool ...
	switch book.GetPrice() {
	case nil:
		fmt.Println("user not set value")
	default:
		fmt.Printf("book.Price: %d\n", book.GetPrice().Value)
	}
}

func main() {

	wrapDemo()
	fieldMaskDemo()
}
