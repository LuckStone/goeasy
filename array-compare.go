package main

import "fmt"

func compare() {
	var del_groups []string
	var add_groups []string

	oldGroup := [...]string{"a", "b", "c"}
	newGroup := [...]string{"b", "c", "d"}

	for _, old_group := range oldGroup {
		match := 0
		for _, new_group := range newGroup {
			if old_group == new_group {
				match = 1
			}
		}
		if match == 0 {
			del_groups = append(del_groups, old_group)
		}
	}

	for _, new_group := range newGroup {
		match := 0
		for _, old_group := range oldGroup {
			if old_group == new_group {
				match = 1
			}
		}
		if match == 0 {
			add_groups = append(add_groups, new_group)
		}
	}

	fmt.Println("delete", del_groups)
	fmt.Println("add", add_groups)

}

func main() {
	fmt.Println("-----------1-------------")
	compare()
}
