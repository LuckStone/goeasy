package main

import (
	"fmt"
)

func main() {
	s := []int{1,2,3}
	s = append(s, 1, 2, 3)
	fmt.Println(s)
}

// [0 0 0 0 0 1 2 3]

// 考点：make默认值和append
// 解答：
// make初始化是由默认值的哦，此处默认值为0
