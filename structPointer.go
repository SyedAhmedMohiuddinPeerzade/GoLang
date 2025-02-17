package main

import (
	"fmt"
)

type Value struct {
	X int
}

func AddWithoutPointerArgument(v Value) { // Strictly NO pointer has to be passed as argument
	v.X = v.X + 10
}

func AddWithPointerArgument(v *Value) { // Strictly Pointer has to be passed as argument
	v.X = v.X + 10
}

func (v Value) AddWithoutPointerReceiver() { // Can be called with or without a Pointer, but value DOESNT gets changed
	v.X = v.X + 10
}

func (v *Value) AddWithPointerReceiver() { // Can be called with or without a Pointer, value gets CHANGED
	v.X = v.X + 10
}

func main() {
	a := Value{5}
	AddWithoutPointerArgument(a)
	fmt.Println(a) // 5

	b := Value{5}
	AddWithPointerArgument(&b)
	fmt.Println(&b) // 15

	c := Value{5}
	c.AddWithoutPointerReceiver()
	fmt.Println(c) // 5

	d := &Value{5}
	d.AddWithPointerReceiver()
	fmt.Println(d) // 15

	e := Value{5}
	e.AddWithPointerReceiver() // We can call without pointer variable.
	fmt.Println(e)             //15

	f := &Value{5}
	f.AddWithoutPointerReceiver() // Value doesnt get changed despite called with pointer variable
	fmt.Println(f)                // 5

	// g := Value{5}
	// AddWithoutPointerArgument(&g) // ERROR: cannot use &g (value of type *Value) as Value value in argument to AddWithoutPointerArgument

	// h := Value{5}
	// AddWithPointerArgument(h) // ERROR:  cannot use h (variable of struct type Value) as *Value value in argument to AddWithPointerArgument
}
