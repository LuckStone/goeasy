package main

import "fmt"

type student struct {
	Name string
	Age  int
}

func pase_student() {
	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	for _, stu := range stus {
		m[stu.Name] = &stu
		fmt.Println(stu.Name)
	}
	fmt.Println(m)
}

func main() {
	pase_student()
}

// 考点：foreach
// 解答：
// 这样的写法初学者经常会遇到的，很危险！ 与Java的foreach一样，都是使用副本的方式。
// 所以m[stu.Name]=&stu实际上一致指向同一个指针，
//  最终该指针的值为遍历的最后一个struct的值拷贝。
// 就像想修改切片元素的属性：

// for _, stu := range stus {
//     stu.Age = stu.Age+10
// }

// 也是不可行的。 大家可以试试打印出来：

// func pase_student() {
//     m := make(map[string]*student)
//     stus := []student{
//         {Name: "zhou", Age: 24},
//         {Name: "li", Age: 23},
//         {Name: "wang", Age: 22},
//     }
//     // 错误写法
//     for _, stu := range stus {
//         m[stu.Name] = &stu
//     }

//     for k,v:=range m{
//         println(k,"=>",v.Name)
//     }

//     // 正确
//     for i:=0;i<len(stus);i++  {
//         m[stus[i].Name] = &stus[i]
//     }
//     for k,v:=range m{
//         println(k,"=>",v.Name)
//     }
// }
