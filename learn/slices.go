package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("s ==", s)

	for i := 0; i < len(s); i++ {
		fmt.Printf("s[%d] == %d\n", i, s[i])
	}
}


/**
 һ�� slice ��ָ��һ�����е�ֵ�����Ұ����˳�����Ϣ��

[]T ��һ��Ԫ������Ϊ T �� slice��

len(s) ���� slice s �ĳ��ȡ� 

*/

// slice ���԰�����������ͣ�������һ�� slice��

func main() {
	// Create a tic-tac-toe board.
	game := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	game[0][0] = "X"
	game[2][2] = "O"
	game[2][0] = "X"
	game[1][0] = "O"
	game[0][2] = "X"

	printBoard(game)
}

func printBoard(s [][]string) {
	for i := 0; i < len(s); i++ {
		fmt.Printf("%s\n", strings.Join(s[i], " "))
	}
}

/**
 slice ����������Ƭ������һ���µ� slice ֵָ����ͬ�����顣

���ʽ s[lo:hi]
��ʾ�� lo �� hi-1 �� slice Ԫ�أ���ǰ�ˣ���������ˡ����

s[lo:lo]�ǿյģ���
s[lo:lo+1]��һ��Ԫ�ء� 
*/

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("s ==", s)
	fmt.Println("s[1:4] ==", s[1:4])

	// ʡ���±����� 0 ��ʼ
	fmt.Println("s[:3] ==", s[:3])

	// ʡ���ϱ���� len(s) ����
	fmt.Println("s[4:] ==", s[4:])
}

//---------------------------------------------------------

func main() {
	a := make([]int, 5)
	printSlice("a", a)   // a len=5 cap=5 [0 0 0 0 0]
	b := make([]int, 0, 5)
	printSlice("b", b)   // b len=0 cap=5 []
	c := b[:2]
	printSlice("c", c)   // c len=2 cap=5 [0 0]
	d := c[2:5]
	printSlice("d", d)   // d len=3 cap=3 [0 0 0]
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

//---------------------------------------------------------

func main() {
	var z []int
	fmt.Println(z, len(z), cap(z))   // [] 0 0
	if z == nil {
		fmt.Println("nil!")          // nil!
	}
}


// ------------------------------------
// ͨ��һ������Ƭ�޸�Ԫ�ػ�Ӱ�쵽ԭʼ��Ƭ�Ķ�ӦԪ��
d := []byte{'r', 'o', 'a', 'd'}
e := d[2:] 
// e == []byte{'a', 'd'}
e[1] = 'm'
// e == []byte{'a', 'm'}
// d == []byte{'r', 'o', 'a', 'm'}




��Ƭ��������copy and append ������

Ҫ������Ƭ���������봴��һ���µġ�������������Ƭ��Ȼ��ԭ����Ƭ�����ݸ��Ƶ��µ���Ƭ�� 

����������һЩ֧�ֶ�̬�������Եĳ���ʵ�֡�

��������ӽ���Ƭ s �����������ȴ���һ��2�� ����������Ƭ t ������ s ��Ԫ�ص� t ��Ȼ�� t ��ֵ�� s ��

t := make([]byte, len(s), (cap(s)+1)*2) // +1 in case cap(s) == 0
for i := range s {
        t[i] = s[i]
}
s = t

ѭ���и��ƵĲ��������� copy ���ú��������copy ������Դ��Ƭ��Ԫ�ظ��Ƶ�Ŀ����Ƭ�� �����ظ���Ԫ�ص���Ŀ��

func copy(dst, src []T) int

copy ����֧�ֲ�ͬ���ȵ���Ƭ֮��ĸ��ƣ���ֻ���ƽ϶���Ƭ�ĳ��ȸ�Ԫ�أ��� ���⣬ copy ����������ȷ����Դ��Ŀ����Ƭ���ص��������

ʹ�� copy ���������ǿ��Լ�����Ĵ���Ƭ�Σ�

t := make([]byte, len(s), (cap(s)+1)*2)
copy(t, s)
s = t

һ�������Ĳ����ǽ�����׷�ӵ���Ƭ��β��������ĺ�����Ԫ��׷�ӵ���Ƭβ���� ��Ҫ�Ļ���������Ƭ����������󷵻ظ��µ���Ƭ��

func AppendByte(slice []byte, data ...byte) []byte {
    m := len(slice)
    n := m + len(data)
    if n > cap(slice) { // if necessary, reallocate
        // allocate double what's needed, for future growth.
        newSlice := make([]byte, (n+1)*2)
        copy(newSlice, slice)
        slice = newSlice
    }
    slice = slice[0:n]
    copy(slice[m:n], data)
    return slice
}

������ AppendByte ��һ���÷���

p := []byte{2, 3, 5}
p = AppendByte(p, 7, 11, 13)
// p == []byte{2, 3, 5, 7, 11, 13}

���� AppendByte �ĺ����Ƚ�ʵ�ã���Ϊ���ṩ����Ƭ������������ȫ���ơ� ���ݳ�����ص㣬����ϣ�������С�Ļ�ϴ�Ŀ飬�����ǳ���ĳ����С�ٷ��䡣

�������������Ҫ��ȫ�Ŀ��ƣ����Go�ṩ��һ�����ú��� append �� ���ڴ�������ϣ����ĺ���ǩ����

func append(s []T, x ...T) []T

append ������ x ׷�ӵ���Ƭ s ��ĩβ�������ڱ�Ҫ��ʱ������������

a := make([]int, 1)
// a == []int{0}
a = append(a, 1, 2, 3)
// a == []int{0, 1, 2, 3}

�����Ҫ��һ����Ƭ׷�ӵ���һ����Ƭβ������Ҫʹ�� ... �﷨����2������չ��Ϊ�����б�

a := []string{"John", "Paul"}
b := []string{"George", "Ringo", "Pete"}
a = append(a, b...) // equivalent to "append(a, b[0], b[1], b[2])"
// a == []string{"John", "Paul", "George", "Ringo", "Pete"}

������Ƭ����ֵ nil ����������һ������Ϊ�����Ƭ�����ǿ�������һ����Ƭ����Ȼ����ѭ�� ������׷�����ݣ�

// Filter returns a new slice holding only
// the elements of s that satisfy f()
func Filter(s []int, fn func(int) bool) []int {
    var p []int // == nil
    for _, v := range s {
        if fn(v) {
            p = append(p, v)
        }
    }
    return p
}

���ܵġ����塱

����ǰ����˵����Ƭ���������Ḵ�Ƶײ�����顣�������齫���������ڴ��У�ֱ�������ٱ����á� ��ʱ����ܻ���Ϊһ��С���ڴ����õ��±������е����ݡ�

���磬 FindDigits �������������ļ����ڴ棬Ȼ��������һ�����������֣����������Ƭ��ʽ���ء�

var digitRegexp = regexp.MustCompile("[0-9]+")

func FindDigits(filename string) []byte {
    b, _ := ioutil.ReadFile(filename)
    return digitRegexp.Find(b)
}

��δ������Ϊ���������ƣ����ص� []byte ָ�򱣴������ļ������顣��Ϊ��Ƭ������ԭʼ�����飬 ���� GC �����ͷ�����Ŀռ䣻ֻ�õ����������ֽ�ȴ���������ļ������ݶ�һֱ�������ڴ��

Ҫ�޸��������⣬���Խ�����Ȥ�����ݸ��Ƶ�һ���µ���Ƭ�У�

func CopyDigits(filename string) []byte {
    b, _ := ioutil.ReadFile(filename)
    b = digitRegexp.Find(b)
    c := make([]byte, len(b))
    copy(c, b)
    return c
}