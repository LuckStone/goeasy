package main

import (
	"fmt"
	"sync"
)

type UserAges struct {
	ages map[string]int
	sync.Mutex
}

func (ua *UserAges) Add(name string, age int) {
	ua.Lock()
	defer ua.Unlock()
	if ua.ages == nil{
		ua.ages = map[string]int{}
	}

	ua.ages[name] = age
}

func (ua *UserAges) Get(name string) int {
	if age, ok := ua.ages[name]; ok {
		return age
	}
	return -1
}

func main() {
	u := UserAges{}
	u.Add("jack", 22)
	age := u.Get("jack")
	fmt.Println(age)
}

// panic: assignment to entry in nil map
// ages没有分配内存

// 考点：map线程安全
// 解答：
// 可能会出现fatal error: concurrent map read and map write. 修改一下看看效果

// func (ua *UserAges) Get(name string) int {
//     ua.Lock()
//     defer ua.Unlock()
//     if age, ok := ua.ages[name]; ok {
//         return age
//     }
//     return -1
// }
