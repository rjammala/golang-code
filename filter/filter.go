package main

import "fmt"
import "container/list"

func main() {
	l := list.New()
	for i := 0; i <= 10; i++ {
		l.PushBack(i)
	}
	r := list.New()
	for e := l.Front(); e != nil; e = e.Next() {
		if e.Value.(int) % 2 == 0 {
			r.PushBack(e.Value)
		}
	}

	for e := r.Front(); e != nil; e = e.Next() {
	    fmt.Println(e.Value)
	}
}
