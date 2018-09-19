package main

// 常量通常会被编译器在预处理阶段直接展开，作为指令数据使用
const cl  = 100

var bl    = 123

func main()  {
    println(&bl,bl)
    println(&cl,cl)
}

// .\test20.go:10: cannot take the address of cl