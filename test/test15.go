package main

type A interface{}

func main(){
    // A类型 就可以接受任何类型了
    var a A
    var i int = 99
    var str string = "wq"
    a = i
    println(a)
    a = str
    println(a)
}