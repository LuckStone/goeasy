package main

import "fmt"



/**
 һ���ṹ�壨 struct ������һ���ֶεļ��ϡ�

���� type �ĺ������������˼������� 
**/

type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{1, 2})
}


// �ṹ���ֶ�ʹ�õ��������
func main() {
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)
}


// �ṹ���ֶο���ͨ���ṹ��ָ�������ʡ�
// ͨ��ָ���ӵķ�����͸���ġ� 
func main() {
	v := Vertex{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(v)
}

// ʹ�� Name: �﷨���Խ��г������ֶΡ����ֶ�����˳���޹ء���
// �����ǰ׺ & ����һ��ָ��ṹ���ָ�롣 
var (
	v1 = Vertex{1, 2}  // ����Ϊ Vertex
	v2 = Vertex{X: 1}  // Y:0 ��ʡ��
	v3 = Vertex{}      // X:0 �� Y:0
	p  = &Vertex{1, 2} // ����Ϊ *Vertex
)

func main() {
	fmt.Println(v1, p, v2, v3)
}

