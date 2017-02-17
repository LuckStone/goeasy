package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := &Vertex{3, 4}
	fmt.Println(v.Abs())
}


/**
 Go û���ࡣȻ������Ȼ�����ڽṹ�������϶��巽����

���������� ������ func �ؼ��ֺͷ�����֮��Ĳ����С� 
*/


type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}


// ��ӡ 1.4142135623730951
/**
 ����Զ԰��е� ���� ���Ͷ������ⷽ����������������Խṹ�塣

���ǣ����ܶ����������������ͻ�������Ͷ��巽����  
*/



// ������������ Vertex �ĸ����������޷��޸�ԭʼֵ��
func (v Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// ������������ Vertex ʵ��
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.Abs())
	v.Scale(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.Abs())
}
