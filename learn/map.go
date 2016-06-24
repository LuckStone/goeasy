package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}
// map ��ʹ��֮ǰ������ make ��������ֵΪ nil �� map �ǿյģ����Ҳ��ܶ��丳ֵ�� 
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

	// ������޸�һ��Ԫ��
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

��� key �� m �У� ok Ϊ true������ ok Ϊ false������ elem �� map ��Ԫ�����͵���ֵ��

ͬ���ģ����� map �ж�ȡĳ�������ڵļ�ʱ������� map ��Ԫ�����͵���ֵ�� 
*/