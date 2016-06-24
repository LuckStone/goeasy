package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}
/**
Go ����ָ�롣 ָ�뱣���˱������ڴ��ַ��

���� *T ��ָ������ T ��ֵ��ָ�롣����ֵ�� nil ��

var p *int

& ���Ż�����һ��ָ�������ö����ָ�롣

i := 42
p = &i

* ���ű�ʾָ��ָ��ĵײ��ֵ��

fmt.Println(*p) // ͨ��ָ�� p ��ȡ i
*p = 21         // ͨ��ָ�� p ���� i

��Ҳ����ͨ����˵�ġ�������á��򡰷�ֱ�����á���

�� C ��ͬ��Go û��ָ�����㡣 
**/