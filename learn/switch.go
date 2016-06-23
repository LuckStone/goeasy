package main

import (
	"fmt"
	"runtime"
	"time"
)

//  switch ���������ϵ��µ�ִ�У���ƥ��ɹ���ʱ��ֹͣ��
//  �����磬
//
//  switch i {
//      case 0:
//      case f():
//  }
//
//  �� i==0 ʱ������� f ���� 

// ������ fallthrough �������������֧���Զ���ֹ
func main() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.", os)
	}
}

// û�������� switch ͬ switch true һ���� 
// ��һ����ʹ�ÿ����ø���������ʽ����д���� if-then-else ���� 
func main() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}