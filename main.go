package main

import (
    "fmt"
    "github.com/sarathmekala/go2/pkg/math"
    "github.com/sarathmekala/go2/pkg/strings"
)

func main() {
    // Use the Add function from the math package
    sum := math.Add(2, 3)
    fmt.Println("Sum:", sum)

    // Use the Concat function from the strings package
    concat := strings.Concat("Hello, ", "World!")
    fmt.Println("Concatenated String:", concat)
}

