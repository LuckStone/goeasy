package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)
}


/**
���� [n]T ��һ���� n ������Ϊ T ��ֵ�����顣

���ʽ

var a [10]int

������� a ��һ����ʮ�����������顣

����ĳ����������͵�һ���֣�������鲻�ܸı��С�� �⿴������һ����Լ�������벻Ҫ���ģ� Go �ṩ�˸��ӱ����ķ�ʽ��ʹ�����顣
*/