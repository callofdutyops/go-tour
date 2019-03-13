package main

import (
	"golang.org/x/tour/tree"
	"fmt"
)

func WalkImpl(t *tree.Tree, ch, quit chan int) {
	if t == nil {
		return
	}
	WalkImpl(t.Left, ch, quit)
	select {
	case ch <- t.Value:
	case <-quit:
		return
	}
	WalkImpl(t.Right, ch, quit)
}

func Walk(t *tree.Tree, ch, quit chan int) {
	WalkImpl(t, ch, quit)
	close(ch)
}

func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)
	quit := make(chan int)
	defer close(quit)

	go Walk(t1, ch1, quit)
	go Walk(t2, ch2, quit)

	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2
		if !ok1 || !ok2 {
			return ok1 == ok2
		}
		if v1 != v2 {
			return false
		}
	}
}

func main() {
	fmt.Println("tree.New(10000) == tree.New(10000)?", Same(tree.New(10000), tree.New(10000)))
	fmt.Println("tree.New(10000) == tree.New(20000)?", Same(tree.New(10000), tree.New(20000)))
}
