package main

func main() {
	i := GetValue()

	switch i.(type) {
	case int:
		println("int")
	case string:
		println("string")
	case interface{}:
		println("interface")
	default:
		println("unknown")
	}

}

func GetValue() int {
	return 1
}

// .\test17.go:6: cannot type switch on non-interface value i (type int)
// type只能使用在interface