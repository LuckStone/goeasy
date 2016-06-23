package main

import (
	"fmt"
	"runtime"
	"time"
)

//  switch 的条件从上到下的执行，当匹配成功的时候停止。
//  （例如，
//
//  switch i {
//      case 0:
//      case f():
//  }
//
//  当 i==0 时不会调用 f 。） 

// 除非以 fallthrough 语句结束，否则分支会自动终止
func main() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.", os)
	}
}

// 没有条件的 switch 同 switch true 一样。 
// 这一构造使得可以用更清晰的形式来编写长的 if-then-else 链。 
func main() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}