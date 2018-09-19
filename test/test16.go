package main

type A interface {}

type B struct {
    name string
}

func main(){
    list := make([]A,3)
    list[0] = 1
    list[1] = "wq"
    list[2] = B{"wq"}
    for _,element := range list {
        // 通过 变量.(类型) 获得返回值 ok 判断 该变量存的是否是 该类型 
        _, ok := element.(B)
        if ok {
            println("is B")
        }
        val2, ok1 := element.(int)
        if ok1 {
            println(val2, "is int")
        }
    }
    //solution 2
    for _,ele := range list {
        // 变量.(type) 仅在switch中 可以使用
        switch ele.(type) {
            case int :
                println("ele is an int")
            case B :
                println("ele is a b")
            default :
                println("ele is default type")
        }
    }
}