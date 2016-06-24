package main

import "fmt"



/**
 一个结构体（ struct ）就是一个字段的集合。

（而 type 的含义跟其字面意思相符。） 
**/

type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{1, 2})
}


// 结构体字段使用点号来访问
func main() {
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)
}


// 结构体字段可以通过结构体指针来访问。
// 通过指针间接的访问是透明的。 
func main() {
	v := Vertex{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(v)
}

// 使用 Name: 语法可以仅列出部分字段。（字段名的顺序无关。）
// 特殊的前缀 & 返回一个指向结构体的指针。 
var (
	v1 = Vertex{1, 2}  // 类型为 Vertex
	v2 = Vertex{X: 1}  // Y:0 被省略
	v3 = Vertex{}      // X:0 和 Y:0
	p  = &Vertex{1, 2} // 类型为 *Vertex
)

func main() {
	fmt.Println(v1, p, v2, v3)
}

