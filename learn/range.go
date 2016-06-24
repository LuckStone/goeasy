package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}


/**
 for ѭ���� range ��ʽ���Զ� slice ���� map ���е���ѭ����

��ʹ�� for ѭ������һ�� slice ʱ��ÿ�ε��� range ����������ֵ�� ��һ���ǵ�ǰ�±꣨��ţ����ڶ����Ǹ��±�����ӦԪ�ص�һ�������� 

*/


func main() {
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}





/**
 ����ͨ����ֵ�� _ ��������ź�ֵ��

���ֻ��Ҫ����ֵ��ȥ�� �� , value �� �Ĳ��ּ��ɡ� 
*/


func main() {
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i)
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}
