package main


// 闭包引用相同变量
func test(x int) (func(),func())  {
    return func() {
        println(x)
        x+=10
    }, func() {
        println(x)
    }
}

func main()  {
    num := 100
    a,b:=test(num)
    a()
    b()
	// num 不会变化
	println(num)
}


//100
//110
//100
