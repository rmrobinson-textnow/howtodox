package main

import (
	"sort"
	"fmt"
)

type SortableThing struct {
	fieldA int
	fieldB string
}

type SortableThings []SortableThing

func (s SortableThings) Len() int {
	return len(s)
}

func (s SortableThings) Less(i, j int) bool {
	return s[i].fieldA < s[j].fieldA
}

func (s SortableThings) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s SortableThings) String() string {
	sort.Sort(s)

	str := "["
	for i, elem := range s {
		if i > 0 {
			str += " "
		}
		str += fmt.Sprintf("%s", elem.fieldB)
	}
	return str + "]"
}

func main() {
	st := SortableThings{
		{
			fieldA: 10,
			fieldB: "ten",
		},
		{
			fieldA: 5,
			fieldB: "five",
		},
		{
			fieldA: 3,
			fieldB: "three",
		},
	}

	fmt.Printf("%s\n", st)
}
