package main

import "fmt"


//  defer �����ӳٺ�����ִ��ֱ���ϲ㺯�����ء�
// �ӳٵ��õĲ������������ɣ��������ϲ㺯������ǰ���������ᱻ���á� 

func main() {
	defer fmt.Println("world")

	fmt.Println("hello")
}




// �ӳٵĺ������ñ�ѹ��һ��ջ�С�����������ʱ�� �ᰴ�պ���ȳ���˳����ñ��ӳٵĺ������á� 
func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
