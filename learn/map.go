package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}
// map 在使用之前必须用 make 来创建；值为 nil 的 map 是空的，并且不能对其赋值。 
var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}


var m = map[string]Vertex{
	"Bell Labs": Vertex{
		40.68433, -74.39967,
	},
	"Google": Vertex{
		37.42202, -122.08408,
	},
}

var m = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}


func main() {
	fmt.Println(m)
}


func main() {
	m := make(map[string]int)

	// 插入或修改一个元素
	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}



/**
elem, ok = m[key]

如果 key 在 m 中， ok 为 true。否则， ok 为 false，并且 elem 是 map 的元素类型的零值。

同样的，当从 map 中读取某个不存在的键时，结果是 map 的元素类型的零值。 
*/