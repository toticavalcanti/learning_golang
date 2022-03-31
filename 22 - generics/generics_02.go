// Starting with version 1.18, Go has added support for
// _generics_, also known as _type parameters_.

package main

import "fmt"

// As an example of a generic function, `MapKeys` takes
// a map of any type and returns a slice of its keys.
// This function has two type parameters - `K` and `V`;
// `K` has the `comparable` _constraint_, meaning that
// we can compare values of this type with the `==` and
// `!=` operators. This is required for map keys in Go.
// `V` has the `any` constraint, meaning that it's not
// restricted in any way (`any` is an alias for `interface{}`).
func MapKeys[K comparable, V any](m map[K]V) []K {
    r := make([]K, 0, len(m))
    for k := range m {
        r = append(r, k)
    }
    return r
}

// As an example of a generic type, `List` is a
// singly-linked list with values of any type.
type List[T any] struct {
    head, tail *element[T]
}

type element[T any] struct {
    next *element[T]
    val  T
}

// We can define methods on generic types just like we
// do on regular types, but we have to keep the type
// parameters in place. The type is `List[T]`, not `List`.
func (lst *List[T]) Push(v T) {
    if lst.tail == nil {
        fmt.Println("TAIL INSIDE IF", lst.tail)
        fmt.Println("HEAD INSIDE IF",lst.head)
        lst.head = &element[T]{val: v}
        fmt.Println("HEAD INSIDE IF AFTER ASSIGNMENT",lst.head)
        lst.tail = lst.head
        fmt.Println("TAIL INSIDE IF AFTER ASSIGNMENT",lst.tail)
    } else {
        fmt.Println()
        fmt.Println("TAIL INSIDE ELSE", lst.tail)
        lst.tail.next = &element[T]{val: v}
        fmt.Println("TAIL-NEXT AFTER ASSIGNMENT", lst.tail.next)
        lst.tail = lst.tail.next
        fmt.Println("TAIL INSIDE ELSE AFTER ASSIGNMENT", lst.tail)
    }
}

func (lst *List[T]) GetAll() []T {
    var elems []T
    for e := lst.head; e != nil; e = e.next {
        elems = append(elems, e.val)
    }
    return elems
}

func main() {
    var m = map[int]string{1: "2", 2: "4", 3: "8"}

    // When invoking generic functions, we can often rely
    // on _type inference_. Note that we don't have to
    // specify the types for `K` and `V` when
    // calling `MapKeys` - the compiler infers them
    // automatically.
    fmt.Println("keys m:", MapKeys(m))

    // ... though we could also specify them explicitly.
    _ = MapKeys[int, string](m)

    lst := List[int]{}
    lst.Push(11) 
    fmt.Println("list:", lst.GetAll())
    lst.Push(13) 
    fmt.Println("list:", lst.GetAll())
    lst.Push(29)
    fmt.Println("list:", lst.GetAll())
    lst.Push(37)
    fmt.Println("list:", lst.GetAll())
    lst.Push(41)

    fmt.Println("list:", lst.GetAll())
}
