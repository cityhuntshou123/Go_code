package main

import "fmt"

func main() {
// 单行注释
/*
	Author by 菜鸟教程
	我是多行注释
*/
fmt.Println("Hello, World!")
fmt.Println("菜鸟教程: runoob.com");

// 字符串连接
fmt.Println("Google" + "Runoob");

// 格式化字符串
var stockcode = 123
var enddate = "2022-2-21"
var url = "Code=%d&endDate=%s"
var target_url = fmt.Sprintf(url, stockcode, enddate)
fmt.Println(target_url)
}