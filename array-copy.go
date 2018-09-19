package main

import "fmt"

func copy1() {
	groups := []string{}

	newGroup := []string{"b", "c", "d"}
	groups = append(groups, newGroup...)
	groups = append(groups, "hello")


	fmt.Println("groups=", groups)
	fmt.Println("newGroup=", newGroup)

}

func copy2() {
	groups := []string{}

	newGroup := []string{"b", "c", "d"}

	groups = newGroup[:]
	groups = append(groups, "hello")

	fmt.Println("groups=", groups)
	fmt.Println("newGroup=", newGroup)
}

func main() {
	copy2()
}
