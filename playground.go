package main

import (
    "fmt"
)

func main() {
    t, z := 0, 1
    t, z = z, t
    fmt.Println(t, z)
}
