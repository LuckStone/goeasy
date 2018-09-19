package main


type Person struct {
    name string
    age int
}

type Student struct {
    // 利用匿名字段可继承该字段类型的方法 同样的可以直接调用
    Person
    no int
    phone string
}

func (p Person) sayhi() {
    println(p.name,p.age)
}

// 重写
func (s Student) sayhi() {
    println(s.name,s.no,s.phone,s.age)
}

func main(){
    s := Student{Person{"wq",20},1,"110"}
    s.sayhi()
}


