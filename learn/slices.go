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
 一个 slice 会指向一个序列的值，并且包含了长度信息。

[]T 是一个元素类型为 T 的 slice。

len(s) 返回 slice s 的长度。 

*/

// slice 可以包含任意的类型，包括另一个 slice。

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
 slice 可以重新切片，创建一个新的 slice 值指向相同的数组。

表达式 s[lo:hi]
表示从 lo 到 hi-1 的 slice 元素，含前端，不包含后端。因此

s[lo:lo]是空的，而
s[lo:lo+1]有一个元素。 
*/

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("s ==", s)
	fmt.Println("s[1:4] ==", s[1:4])

	// 省略下标代表从 0 开始
	fmt.Println("s[:3] ==", s[:3])

	// 省略上标代表到 len(s) 结束
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
// 通过一个新切片修改元素会影响到原始切片的对应元素
d := []byte{'r', 'o', 'a', 'd'}
e := d[2:] 
// e == []byte{'a', 'd'}
e[1] = 'm'
// e == []byte{'a', 'm'}
// d == []byte{'r', 'o', 'a', 'm'}




切片的生长（copy and append 函数）

要增加切片的容量必须创建一个新的、更大容量的切片，然后将原有切片的内容复制到新的切片。 

整个技术是一些支持动态数组语言的常见实现。

下面的例子将切片 s 容量翻倍，先创建一个2倍 容量的新切片 t ，复制 s 的元素到 t ，然后将 t 赋值给 s ：

t := make([]byte, len(s), (cap(s)+1)*2) // +1 in case cap(s) == 0
for i := range s {
        t[i] = s[i]
}
s = t

循环中复制的操作可以由 copy 内置函数替代。copy 函数将源切片的元素复制到目的切片。 它返回复制元素的数目。

func copy(dst, src []T) int

copy 函数支持不同长度的切片之间的复制（它只复制较短切片的长度个元素）。 此外， copy 函数可以正确处理源和目的切片有重叠的情况。

使用 copy 函数，我们可以简化上面的代码片段：

t := make([]byte, len(s), (cap(s)+1)*2)
copy(t, s)
s = t

一个常见的操作是将数据追加到切片的尾部。下面的函数将元素追加到切片尾部， 必要的话会增加切片的容量，最后返回更新的切片：

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

下面是 AppendByte 的一种用法：

p := []byte{2, 3, 5}
p = AppendByte(p, 7, 11, 13)
// p == []byte{2, 3, 5, 7, 11, 13}

类似 AppendByte 的函数比较实用，因为它提供了切片容量增长的完全控制。 根据程序的特点，可能希望分配较小的活较大的块，或则是超过某个大小再分配。

但大多数程序不需要完全的控制，因此Go提供了一个内置函数 append ， 用于大多数场合；它的函数签名：

func append(s []T, x ...T) []T

append 函数将 x 追加到切片 s 的末尾，并且在必要的时候增加容量。

a := make([]int, 1)
// a == []int{0}
a = append(a, 1, 2, 3)
// a == []int{0, 1, 2, 3}

如果是要将一个切片追加到另一个切片尾部，需要使用 ... 语法将第2个参数展开为参数列表。

a := []string{"John", "Paul"}
b := []string{"George", "Ringo", "Pete"}
a = append(a, b...) // equivalent to "append(a, b[0], b[1], b[2])"
// a == []string{"John", "Paul", "George", "Ringo", "Pete"}

由于切片的零值 nil 用起来就像一个长度为零的切片，我们可以声明一个切片变量然后在循环 中向它追加数据：

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

可能的“陷阱”

正如前面所说，切片操作并不会复制底层的数组。整个数组将被保存在内存中，直到它不再被引用。 有时候可能会因为一个小的内存引用导致保存所有的数据。

例如， FindDigits 函数加载整个文件到内存，然后搜索第一个连续的数字，最后结果以切片方式返回。

var digitRegexp = regexp.MustCompile("[0-9]+")

func FindDigits(filename string) []byte {
    b, _ := ioutil.ReadFile(filename)
    return digitRegexp.Find(b)
}

这段代码的行为和描述类似，返回的 []byte 指向保存整个文件的数组。因为切片引用了原始的数组， 导致 GC 不能释放数组的空间；只用到少数几个字节却导致整个文件的内容都一直保存在内存里。

要修复整个问题，可以将感兴趣的数据复制到一个新的切片中：

func CopyDigits(filename string) []byte {
    b, _ := ioutil.ReadFile(filename)
    b = digitRegexp.Find(b)
    c := make([]byte, len(b))
    copy(c, b)
    return c
}