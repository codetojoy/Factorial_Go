
package main

import (
    "fmt"

    "github.com/codetojoy/config"
    "github.com/codetojoy/factorial"
)

func process(n int) {
    config := config.New(n)

    for i:= 2; i <= n; i++ {
        f := factorial.NewFactorial(i, config)
        f.Compute()
        fmt.Printf("TRACER f: %d -> %v \n", i, f.GetPureFactorsString())
    }
}

func main() {
    n := 10
    process(n)

    fmt.Println("Ready.")
}
