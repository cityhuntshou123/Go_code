package main

import "fmt"

// 常量-代码示例
func main() {
	
	// 显式类型定义
	const b string = "abc"
	
	// 隐式类型定义
	const c = "abc"
	
	const LENGTH int = 10;
	const WIDTH int = 5;
	var area int;
	const a, b, c = 1, false,  "str"  // 多重赋值
	
	area = LENGTH * WIDTH
	fmt.Printf("面积为: %d", area) // 50 
	println() 
	println(a, b, c) // 1, false, "abc"
	
}