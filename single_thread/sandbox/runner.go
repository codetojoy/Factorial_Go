
package main

import (
    "fmt"
    "./foo"
    "./bar"
)

func main() {
    myIndex := New(4140)
    fmt.Printf("TRACER myIndex: %v\n", myIndex.String())

    fooIndex := foo.New(5150)
    fmt.Printf("TRACER fooIndex: %v\n", fooIndex.String())

    barUtility := bar.New(6160)
    fmt.Printf("TRACER barUtility: %v\n", barUtility.String())

    fmt.Println("Ready.")
}
