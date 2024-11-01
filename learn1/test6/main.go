package main

import (
	"fmt"
	note "micro/learn1/test6/note"
)

// noeof 使用示例 :  解决定义多个字段但是只能选择其中一个使用的场景

func oneofDemo() {

	// client: 指定数据
	// 下面演示的 oneof 示例中只选择其中一个字段填充，这里选择 email
	ret := &note.Notice{
		Msg: "welcome to boliang",
		NoticeWay: &note.Notice_Email{
			Email: "1964475295@qq.com",
		},
	}

	// server: 类型断言
	// 判断 ret.NoticeWay 到底是哪种类型，然后执行对应的操作即可
	switch v := ret.NoticeWay.(type) {
	case *note.Notice_Email:
		EmailWay(v)
	case *note.Notice_Wechat:
		WeChatWay(v)
	case *note.Notice_Phone:
		PhoneWay(v)
	}
}

func EmailWay(in *note.Notice_Email) {
	fmt.Printf("client use email way : %v\n", in.Email)
}

func WeChatWay(in *note.Notice_Wechat) {
	fmt.Printf("client use wechat way: %v\n", in.Wechat)
}

func PhoneWay(in *note.Notice_Phone) {
	fmt.Printf("client use phone way: %v\n", in.Phone)
}

func main() {
	oneofDemo()
}
