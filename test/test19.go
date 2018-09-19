package main


func main() {

	println(DeferFunc1(1))
	println(DeferFunc2(1))
	println(DeferFunc3(1))
}


// 函数返回值t作用域为整个函数，在return之前defer会被执行，所以t会被修改，返回4
func DeferFunc1(i int) (t int) {
	t = i
	defer func() {
		t += 3
	}()
	return t
}

// t的作用域为函数，返回1
func DeferFunc2(i int) int {
	t := i
	defer func() {
		t += 3
	}()
	return t
}

// 返回3
func DeferFunc3(i int) (t int) {
	defer func() {
		t += i
	}()
	return 2
}

// 需要明确一点是defer需要在函数结束前执行。 
// 函数返回值名字会在函数起始处被初始化为对应类型的零值并且作用域为整个函数