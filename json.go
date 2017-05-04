package main

import "encoding/json"
import "fmt"
import "os"

type Response1 struct {
    Page   int
    Fruits []string
}

type Response2 struct {
    Page   int      `json:"page"`
    Fruits []string `json:"fruits"`
}

type CommonArg struct {
    SessionId int64  `json:",string"`
    Op        string `json:"Op,omitempty"`
    AppId     string `json:"appId,omitempty"`
    Online    bool
}

func main() {
    bolB, _ := json.Marshal(true)
    fmt.Println(string(bolB)) // true

    intB, _ := json.Marshal(1)
    fmt.Println(string(intB)) // 1

    fltB, _ := json.Marshal(2.34)
    fmt.Println(string(fltB)) // 2.34

    strB, _ := json.Marshal("gopher")
    fmt.Println(string(strB)) // "gopher"

    slcD := []string{"apple", "peach", "pear"}
    slcB, _ := json.Marshal(slcD)
    fmt.Println(string(slcB)) //  ["apple","peach","pear"]

    mapD := map[string]int{"apple": 5, "lettuce": 7}
    mapB, _ := json.Marshal(mapD)
    fmt.Println(string(mapB)) // {"apple":5,"lettuce":7}

    res1D := &Response1{
        Page:   1,
        Fruits: []string{"apple", "peach", "pear"}}
    res1B, _ := json.Marshal(res1D)
    fmt.Println(string(res1B)) // {"Page":1,"Fruits":["apple","peach","pear"]}

    res2D := &Response2{
        Page:   1,
        Fruits: []string{"apple", "peach", "pear"}}
    res2B, _ := json.Marshal(res2D)
    fmt.Println(string(res2B)) // {"Page":1,"Fruits":["apple","peach","pear"]}

    aaa := &CommonArg{
        SessionId: 12345,
        Op:        "werdfsdde22233",
        AppId:     "654343dfddd33424dd",
        Online:    false,
    }
    aaa1, _ := json.Marshal(aaa)
    fmt.Println(string(aaa1)) // {"SessionId":"12345","Op":"werdfsdde22233","appId":"654343dfddd33424dd","Online":false}

    byt := []byte(`{"num":6.0,"strs":["a","b"]}`)

    var dat map[string]interface{}

    if err := json.Unmarshal(byt, &dat); err != nil {
        panic(err)
    }
    fmt.Println(dat) // map[strs:[a b] num:6]

    num := dat["num"].(float64)
    fmt.Println(num) // 6

    strs := dat["strs"].([]interface{})
    str1 := strs[0].(string)
    fmt.Println(str1) // a

    str := `{"page": 1, "fruits": ["apple", "peach"]}`
    res := &Response2{}
    json.Unmarshal([]byte(str), &res)
    fmt.Println(res)           // &{1 [apple peach]}
    fmt.Println(res.Fruits[0]) // apple

    enc := json.NewEncoder(os.Stdout)
    d := map[string]int{"apple": 5, "lettuce": 7}
    enc.Encode(d) // {"apple":5,"lettuce":7}
}

 

注意，这里 json的 struct field's tag 规范定义如下：

在Go语言里，StructTag是一个标记字符串，此字符串可跟随在Struct中字段定义的后面。

StructTag就是一系列的 key:”value” 形式的组合，其中key是一个不可为空的字符串，key-value组合可以有多个，空格分隔。

StructTag有什么用？！StructTag主要解决了不同类型数据集合间(Struct,Json,Table等)转换中键值Key定义不一样的问题。StructTag可以理解为一个不同数据类型键值Key的映射表Map, 在StructTag中可以定义不同数据集合键值和Struct中Key值的映射关系，这样方便了Struct数据转为其他类型数据的过程。

在StructTag中加入”omitempty”, 标识该字段的数据可忽略。
- 指定到一个field时，无论有没有值将Person序列化成json时都会忽略该字段

参考下面代码， 代码来自：http://studygolang.com/articles/1698：

 

package main

import (
    "encoding/json"
    "fmt"
)

//tag中的第一个参数是用来指定别名
//比如Name 指定别名为 username `json:"username"`
//如果不想指定别名但是想指定其他参数用逗号来分隔
//omitempty 指定到一个field时 如果在赋值时对该属性赋值 或者 对该属性赋值为 zero value
//那么将Person序列化成json时会忽略该字段
//- 指定到一个field时
//无论有没有值将Person序列化成json时都会忽略该字段
//string 指定到一个field时
//比如Person中的Count为int类型 如果没有任何指定在序列化
//到json之后也是int 比如这个样子 "Count":0
//但是如果指定了string之后序列化之后也是string类型的
//那么就是这个样子"Count":"0"
type Person struct {
    Name        string `json:"username"`
    Age         int
    Gender      bool `json:",omitempty"`
    Profile     string
    OmitContent string `json:"-"`
    Count       int    `json:",string"`
}

func main() {

    var p *Person = &Person{
        Name:        "brainwu",
        Age:         21,
        Gender:      true,
        Profile:     "I am ghj1976",
        OmitContent: "OmitConent",
    }
    if bs, err := json.Marshal(&p); err != nil {
        panic(err)
    } else {
        //result --> {"username":"brainwu","Age":21,"Gender":true,"Profile":"I am ghj1976","Count":"0"}
        fmt.Println(string(bs))
    }

    var p2 *Person = &Person{
        Name:        "brainwu",
        Age:         21,
        Profile:     "I am ghj1976",
        OmitContent: "OmitConent",
    }
    if bs, err := json.Marshal(&p2); err != nil {
        panic(err)
    } else {
        //result --> {"username":"brainwu","Age":21,"Profile":"I am ghj1976","Count":"0"}
        fmt.Println(string(bs))
    }

    // slice 序列化为json
    var aStr []string = []string{"Go", "Java", "Python", "Android"}
    if bs, err := json.Marshal(aStr); err != nil {
        panic(err)
    } else {
        //result --> ["Go","Java","Python","Android"]
        fmt.Println(string(bs))
    }

    //map 序列化为json
    var m map[string]string = make(map[string]string)
    m["Go"] = "No.1"
    m["Java"] = "No.2"
    m["C"] = "No.3"
    if bs, err := json.Marshal(m); err != nil {
        panic(err)
    } else {
        //result --> {"C":"No.3","Go":"No.1","Java":"No.2"}
        fmt.Println(string(bs))
    }
}