package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

// ������֮ǰִ��һ�������
// �������䶨��ı�������������� if ��Χ֮��
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

// �� if �ı����䶨��ı���ͬ���������κζ�Ӧ�� else ����ʹ�á� 
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// ���￪ʼ�Ͳ���ʹ�� v ��
	return lim
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))
}