package main

import "fmt"

type firstType struct {
	fieldA string
	fieldB int
}

type secondType struct {
	fieldOne int
	fieldTwo float64
}

func typeswitch(iface interface{}) {
	// Here we switch on the type of iface
	// We assign the value to t, which allows us to access the relevant fields in the scope of that case.
	switch t := iface.(type) {
	case firstType:
		fmt.Printf("fieldA: %s, fieldB: %d\n", t.fieldA, t.fieldB)
	case secondType:
		fmt.Printf("fieldOne: %d, fieldTwo: %0.2f\n", t.fieldOne, t.fieldTwo)
	}
}

func typecheck(iface interface{}) {
	// Here we check if we can cast the interface to the specified type (firstType)
	// If the cast is okay, we are able to access the fields in a type-safe manner.
	// If the cast wasn't okay, we don't do anything.
	if ft, ok := iface.(firstType); ok {
		fmt.Printf("fieldA: %s, fieldB: %d\n", ft.fieldA, ft.fieldB)
	}

}

func main() {
	ft := firstType{
		fieldA: "string in field A",
		fieldB: 42,
	}

	st := secondType{
		fieldOne: 31,
		fieldTwo: 3.141592,
	}

	typeswitch(ft)

	typeswitch(st)

	typecheck(ft)

	typecheck(st)
}
