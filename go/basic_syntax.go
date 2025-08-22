package main

import (
	"fmt"
	"time"
)

// 1. 变量声明和赋值
func variableExamples() {
	fmt.Println("=== 变量声明和赋值 ===")

	// 方式1：var 关键字声明
	var name string = "张三"
	var age int = 25

	// 方式2：短变量声明（推荐）
	height := 175.5
	isStudent := true

	// 方式3：批量声明
	var (
		city = "北京"
	)

	// 方式4：类型推断
	var score = 95.5

	fmt.Printf("姓名: %s, 年龄: %d, 身高: %.1f\n", name, age, height)
	fmt.Printf("是否学生: %t, 城市: %s, 分数: %.1f\n", isStudent, city, score)
}

// 2. 常量声明
func constantExamples() {
	fmt.Println("\n=== 常量声明 ===")

	// 单个常量
	const PI = 3.14159

	// 批量常量
	const (
		SUNDAY    = 0
		MONDAY    = 1
		TUESDAY   = 2
		WEDNESDAY = 3
	)

	// iota 枚举（自动递增）
	const (
		RED   = iota // 0
		GREEN        // 1
		BLUE         // 2
	)

	fmt.Printf("PI: %.5f\n", PI)
	fmt.Printf("颜色枚举: RED=%d, GREEN=%d, BLUE=%d\n", RED, GREEN, BLUE)
}

// 3. 基本数据类型
func dataTypeExamples() {
	fmt.Println("\n=== 基本数据类型 ===")

	// 整数类型
	var intNum int = 42
	var int8Num int8 = 127
	var int64Num int64 = 9223372036854775807

	// 浮点数
	var float32Num float32 = 3.14
	var float64Num float64 = 3.14159265359

	// 字符串
	var str string = "Hello, Go!"

	// 布尔值
	var boolVal bool = true

	// 复数
	var complexNum complex128 = 3 + 4i

	fmt.Printf("整数: %d, %d, %d\n", intNum, int8Num, int64Num)
	fmt.Printf("浮点数: %.2f, %.10f\n", float32Num, float64Num)
	fmt.Printf("字符串: %s\n", str)
	fmt.Printf("布尔值: %t\n", boolVal)
	fmt.Printf("复数: %v\n", complexNum)
}

// 4. 控制结构
func controlStructureExamples() {
	fmt.Println("\n=== 控制结构 ===")

	// if 条件语句
	age := 18
	if age >= 18 {
		fmt.Println("成年人")
	} else {
		fmt.Println("未成年人")
	}

	// if 语句可以包含初始化
	if score := 85; score >= 90 {
		fmt.Println("优秀")
	} else if score >= 80 {
		fmt.Println("良好")
	} else {
		fmt.Println("需要努力")
	}

	// for 循环
	fmt.Println("for 循环:")
	for i := 0; i < 5; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	// range 循环（类似 foreach）
	fruits := []string{"苹果", "香蕉", "橙子"}
	fmt.Println("range 循环:")
	for index, fruit := range fruits {
		fmt.Printf("索引 %d: %s\n", index, fruit)
	}

	// switch 语句
	day := 3
	switch day {
	case 1:
		fmt.Println("星期一")
	case 2:
		fmt.Println("星期二")
	case 3:
		fmt.Println("星期三")
	default:
		fmt.Println("其他日子")
	}
}

// 5. 函数定义和调用
func functionExamples() {
	fmt.Println("\n=== 函数定义和调用 ===")

	// 基本函数
	result := add(10, 20)
	fmt.Printf("10 + 20 = %d\n", result)

	// 多返回值函数
	quotient, remainder := divide(17, 5)
	fmt.Printf("17 ÷ 5 = %d 余 %d\n", quotient, remainder)

	// 命名返回值
	area, perimeter := rectangle(5, 3)
	fmt.Printf("矩形面积: %.1f, 周长: %.1f\n", area, perimeter)
}

// 基本函数
func add(a, b int) int {
	return a + b
}

// 多返回值函数
func divide(a, b int) (int, int) {
	return a / b, a % b
}

// 命名返回值
func rectangle(length, width float64) (area float64, perimeter float64) {
	area = length * width
	perimeter = 2 * (length + width)
	return // 裸返回，返回命名返回值
}

// 6. 错误处理
func errorHandlingExamples() {
	fmt.Println("\n=== 错误处理 ===")

	// 模拟一个可能出错的函数
	result, err := safeDivide(10, 2)
	if err != nil {
		fmt.Printf("错误: %v\n", err)
	} else {
		fmt.Printf("结果: %.2f\n", result)
	}

	// 除零错误
	result, err = safeDivide(10, 0)
	if err != nil {
		fmt.Printf("错误: %v\n", err)
	} else {
		fmt.Printf("结果: %.2f\n", result)
	}
}

func safeDivide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("除数不能为零")
	}
	return a / b, nil
}

// 7. 切片和映射
func sliceAndMapExamples() {
	fmt.Println("\n=== 切片和映射 ===")

	// 切片（动态数组）
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Printf("切片: %v, 长度: %d, 容量: %d\n", numbers, len(numbers), cap(numbers))

	// 添加元素
	numbers = append(numbers, 6, 7)
	fmt.Printf("添加后: %v\n", numbers)

	// 切片操作
	subSlice := numbers[1:4] // 索引1到3（不包含4）
	fmt.Printf("子切片: %v\n", subSlice)

	// 映射（字典）
	person := map[string]interface{}{
		"name": "李四",
		"age":  30,
		"city": "上海",
	}

	fmt.Printf("映射: %v\n", person)
	fmt.Printf("姓名: %s\n", person["name"])

	// 检查键是否存在
	if age, exists := person["age"]; exists {
		fmt.Printf("年龄: %v\n", age)
	}
}

// 8. 结构体
func structExamples() {
	fmt.Println("\n=== 结构体 ===")

	// 定义结构体
	type Person struct {
		Name string
		Age  int
		City string
	}

	// 创建结构体实例
	person1 := Person{
		Name: "王五",
		Age:  28,
		City: "深圳",
	}

	// 访问字段
	fmt.Printf("人员信息: %s, %d岁, 来自%s\n", person1.Name, person1.Age, person1.City)
}

// 9. goroutine 和 channel
func concurrencyExamples() {
	fmt.Println("\n=== 并发编程 ===")

	// 创建 channel
	ch := make(chan string)

	// 启动 goroutine
	go func() {
		time.Sleep(1 * time.Second)
		ch <- "Hello from goroutine!"
	}()

	// 从 channel 接收数据
	message := <-ch
	fmt.Printf("收到消息: %s\n", message)

	// 带缓冲的 channel
	bufferedCh := make(chan int, 3)
	bufferedCh <- 1
	bufferedCh <- 2
	bufferedCh <- 3

	fmt.Printf("缓冲channel长度: %d\n", len(bufferedCh))

	// 读取所有数据
	for i := 0; i < 3; i++ {
		fmt.Printf("读取: %d\n", <-bufferedCh)
	}
}

func main() {
	fmt.Println("Go语言基础语法示例")
	fmt.Println("==================")

	variableExamples()
	constantExamples()
	dataTypeExamples()
	controlStructureExamples()
	functionExamples()
	errorHandlingExamples()
	sliceAndMapExamples()
	structExamples()
	concurrencyExamples()

	fmt.Println("\n==================")
	fmt.Println("基础语法示例完成")
}
