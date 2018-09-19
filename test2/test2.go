package main



func main() {
    pase_student()
}

type student struct {
    Name string
    Age  int
}

func pase_student() {
    m := make(map[string]student)
    stus := []student{
        {Name: "zhou", Age: 24},
        {Name: "li", Age: 23},
        {Name: "wang", Age: 22},
    }
	for _, stu := range stus {
		stu.Age = stu.Age+10
	}
	for _,v:=range stus{
        println(v.Name,"=>", v.Age)
    }

    for _, stu := range stus {
        m[stu.Name] = stu
    }
	
	// v是一个副本，执行之后对m没有影响
	for _,v:=range m{
        v.Age = v.Age + 100
    }

	for k,v:=range m{
        println(k,"=>",v.Name,"age=>",v.Age)
    }
}