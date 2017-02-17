package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}

//-----------------------------------------------------


type ErrNegativeSqrt float64


//*ע�⣺* �� Error �����ڵ��� fmt.Sprint(e) �����ó���������ѭ��������ͨ����ת�� e ������������⣺fmt.Sprint(float64(e))����˼������Ϊʲô�أ� 
func (e ErrNegativeSqrt) Error() string{
	return fmt.Sprintf("cannot Sqrt negative number:%v", float64(e))
}


func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	return 0, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}