package main

import "fmt"

func main() {
	var a *int
	var a_array_int []int
	var a_map map[string] int
	var a_chan chan int
	var a_func func(string) int
	var a_error error // error 是接口
	
	fmt.Println(a)
	fmt.Println(a_array_int)
	fmt.Println(a_map)
	fmt.Println(a_chan)
	fmt.Println(a_func)
	fmt.Println(a_error)
	
	var i int // 0
	var f float64 // 0
	var b bool // false
	var s string // ""
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
	
	// 根据值自行判定变量类型
	var d = true
	fmt.Println(d)
	
	var intVal int
	// intVal := 1
	intVal1 := 1
	fmt.Println(intVal)
	fmt.Println(intVal1)
	
	// var f string = "Runoob" => f := "Runoob"
	f_str :=  "Runoob"
	fmt.Println(f_str)
}