package main

import "fmt"

//  option 选项模式
const DefaultAge = 18
const defaultName = "chen"
const defaultScore = 80
const defaultWeight = 50.5

// 定义一个正常的结构体
type Config struct {
	Name  string
	Age   int
	Hobby string

	Level struct{}
	Money Info
}

type Info struct {
	Profit float64
}

// 正常结构体的初始化，无法传入默认参数
func NewConfig(name, hobby string, age int) *Config {

	return &Config{
		Name:  name,
		Age:   age,
		Hobby: hobby,
	}
}

// 目的： 想要 name , hobby必须传，age 可传可不传
// 不方便结构体字段的拓展，比如说结构体内部由自定义的字段，并且实现可传可不传，这时候，下面的代码就相当于是写死了
func NewConfig2(name string, hobby string, age ...int) *Config {

	valueAge := DefaultAge
	if len(age) > 0 {
		valueAge = age[0]
	}
	return &Config{
		Name:  name,
		Age:   valueAge,
		Hobby: hobby,
	}
}

// ----------- option基础 -----------
// 闭包 + 可变长参数

// 定义一个函数类型
type FuncConfigOption func(*Config)

func NewConfig3(name string, hobby string, opts ...FuncConfigOption) *Config {

	config := &Config{
		Name:  name,
		Hobby: hobby,
		Age:   DefaultAge,
	}

	// 针对可能传进来的 FuncConfigOption 做参数处理
	for _, opt := range opts {
		opt(config)
	}
	return config
}

// 闭包结构实现默认参数处理
// 通常函数名字使用 [With + Value_Name]
func WithAge(age int) FuncConfigOption {

	return func(cfg *Config) {
		cfg.Age = age
	}
}

// 结构体内部嵌套的结构体也是同样
func WithMoney(profit float64) FuncConfigOption {

	return func(cfg *Config) {
		cfg.Money.Profit = profit
	}
}

func main() {

	cfg := NewConfig3("rabbit", "dance")
	fmt.Printf("cfg:%+v\n", cfg)

	cfg = NewConfig3("rabbit", "chen", WithAge(100))
	fmt.Printf("cfg:%+v\n", cfg)

	cfg = NewConfig3("rabbit", "chen", WithAge(100), WithMoney(45.92))
	fmt.Printf("cfg:%+v\n", cfg)

	cfg.Age = 50 // 外部包直接修改，你奈我何? ? ! ! (●'◡'●)

	fmt.Println("-----------------------------")

	work := NewWorker(18)
	fmt.Printf("worker:%+v\n", work)

	work = NewWorker(18, WithWorkerName("rabbit"), WithWorkerScore(99))
	fmt.Printf("worker:%+v\n", work)
	// 这样定义的结构体对象在包外可以隐藏内部的具体实现
	// 包内仍然可以修改

}

// ---------- option 进阶 ----------
// 使用接口类型去隐藏具体实现，结构体名和字段名使用小写，防止在其他的包中可以修改

// 1. 定义一个只对包内使用的结构体
type worker struct {
	name   string
	age    int
	score  int
	weight float64
}

// 2. 定义一个接口
type WorkerOption interface {
	apply(*worker)
}

// 3. 定义一个函数结构体
type funcOption struct {
	f func(*worker)
}

// 4. 函数结构体实现接口方法
func (f funcOption) apply(work *worker) {
	f.f(work)
}

// 5. 函数结构体的实现
func NewFuncOption(f func(*worker)) funcOption {
	return funcOption{f: f}
}

// 6. 真正的实现默认字段名(调用的时候可选择是否传入)
// 需要传入的字段每个都定义一个 With 函数吗 ? ? ?
func WithWorkerName(name string) WorkerOption {

	return NewFuncOption(func(work *worker) {
		work.name = name
	})
}

func WithWorkerScore(score int) WorkerOption {

	return NewFuncOption(func(work *worker) {
		work.score = score
	})
}

// 7. 对外实现用户调用(携带默认参数)
func NewWorker(age int, opts ...WorkerOption) *worker {

	// 这里传入所有的默认值
	// 需要自定义的字的，每个字段定义一个WithXxxx函数即可
	work := &worker{
		name:   defaultName,
		age:    age,
		score:  defaultScore,
		weight: defaultWeight,
	}
	// 上面的代码相当于把上面配置好了之后再去更改默认的值...

	for _, opt := range opts {
		opt.apply(work)
	}
	return work
}
