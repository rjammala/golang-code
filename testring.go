package main

import "fmt"
import "github.com/rjammala/golang-ring"

func main() {
	r := &ring.Ring{}
	r.SetCapacity(10)
	for i := 0; i < 100; i++ {
		r.Enqueue(i)
	}
	n := 5
	for _, r := range r.Values() {
		if n > 0 {
			fmt.Println(r)
			n--
		}
	}
}
