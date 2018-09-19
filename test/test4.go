package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}
	wg.Add(20)
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("i=", i)
			wg.Done()
		}()
	}
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("i:", i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

// i: 9
// i= 10
// i= 10
// i= 10
// i= 10
// i= 10
// i= 10
// i= 10
// i= 10
// i= 10
// i= 10
// i: 0
// i: 1
// i: 2
// i: 3
// i: 4
// i: 5
// i: 6
// i: 7
// i: 8

// 考点：go执行的随机性和闭包
// 解答：
// 谁也不知道执行后打印的顺序是什么样的，所以只能说是随机数字。 但是A:均为输出10，B:从0~9输出(顺序不定)。
// 第一个go func中i是外部for的一个变量，地址不变化。遍历完成后，最终i=10。 故go func执行时，i的值始终是10。

// 第二个go func中i是函数参数，与外部for中的i完全是两个变量。 尾部(i)将发生值拷贝，go func内部指向值拷贝地址。
