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

 

ע�⣬���� json�� struct field's tag �淶�������£�

��Go�����StructTag��һ������ַ��������ַ����ɸ�����Struct���ֶζ���ĺ��档

StructTag����һϵ�е� key:��value�� ��ʽ����ϣ�����key��һ������Ϊ�յ��ַ�����key-value��Ͽ����ж�����ո�ָ���

StructTag��ʲô�ã���StructTag��Ҫ����˲�ͬ�������ݼ��ϼ�(Struct,Json,Table��)ת���м�ֵKey���岻һ�������⡣StructTag�������Ϊһ����ͬ�������ͼ�ֵKey��ӳ���Map, ��StructTag�п��Զ��岻ͬ���ݼ��ϼ�ֵ��Struct��Keyֵ��ӳ���ϵ������������Struct����תΪ�����������ݵĹ��̡�

��StructTag�м��롱omitempty��, ��ʶ���ֶε����ݿɺ��ԡ�
- ָ����һ��fieldʱ��������û��ֵ��Person���л���jsonʱ������Ը��ֶ�

�ο�������룬 �������ԣ�http://studygolang.com/articles/1698��

 

package main

import (
    "encoding/json"
    "fmt"
)

//tag�еĵ�һ������������ָ������
//����Name ָ������Ϊ username `json:"username"`
//�������ָ������������ָ�����������ö������ָ�
//omitempty ָ����һ��fieldʱ ����ڸ�ֵʱ�Ը����Ը�ֵ ���� �Ը����Ը�ֵΪ zero value
//��ô��Person���л���jsonʱ����Ը��ֶ�
//- ָ����һ��fieldʱ
//������û��ֵ��Person���л���jsonʱ������Ը��ֶ�
//string ָ����һ��fieldʱ
//����Person�е�CountΪint���� ���û���κ�ָ�������л�
//��json֮��Ҳ��int ����������� "Count":0
//�������ָ����string֮�����л�֮��Ҳ��string���͵�
//��ô�����������"Count":"0"
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

    // slice ���л�Ϊjson
    var aStr []string = []string{"Go", "Java", "Python", "Android"}
    if bs, err := json.Marshal(aStr); err != nil {
        panic(err)
    } else {
        //result --> ["Go","Java","Python","Android"]
        fmt.Println(string(bs))
    }

    //map ���л�Ϊjson
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