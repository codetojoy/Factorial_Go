
package main

import (
    "fmt"
    "log"
    "os"

    "github.com/codetojoy/config"
    "github.com/codetojoy/factorial"
    "github.com/codetojoy/util"
)

func processN(n int) {
    config := config.New(n)

    for i:= 2; i <= n; i++ {
        f := factorial.NewFactorial(i, config)
        f.Compute()
        fmt.Printf("TRACER f: %d -> %v \n", i, f.GetPureFactorsString())
    }
}

func searchChunk(chunkLow int, chunkHigh int, factorialIndex *factorial.FactorialIndex) {
    for a := chunkLow; a <= chunkHigh; a++ {
        factorialA := factorialIndex.Get(a)
        for b := 2; b <= a; b++ {
            factorialB := factorialIndex.Get(b)
            for c := 2; c <= a; c++ {
                factorialC := factorialIndex.Get(c)
                result, e := factorialA.IsProductMatch(factorialB, factorialC)

                if e != nil {
                    log.Fatal(e)
                } else if result {
                    fmt.Printf("TRACER %d! = %d! x %d!\n", a, b, c)
                }
            }
        }
    }
}

func search(n int) {
    config := config.New(n)
    factorialIndex := factorial.NewFactorialIndex(config)
    isDone := false
    chunkIndex := 0
    const chunkSize = 20

    for ! isDone {
        chunkLow, chunkHigh := util.GetChunk(chunkIndex, chunkSize, n)
        searchChunk(chunkLow, chunkHigh, factorialIndex)

        if chunkHigh >= n {
            isDone = true
        }

        chunkIndex++
    }
}

func main() {
    numArgs := len(os.Args)

    if numArgs > 1 {
        mode := os.Args[1]

        if mode == "process50" {
            n := 50
            processN(n)
        } else if mode == "search" {
            x := 121
            search(x)
        } else {
            fmt.Println("ERROR: unknown mode")
            os.Exit(1)
        }
    }

    fmt.Println("Ready.")
}
