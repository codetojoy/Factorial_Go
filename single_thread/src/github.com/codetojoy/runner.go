
package main

import (
    "fmt"

    "github.com/codetojoy/config"
    "github.com/codetojoy/factor"
    "github.com/codetojoy/factorial"
    "github.com/codetojoy/prime"
)

func process1(n int) {
    primeIndex := prime.New(n)

    max := 10

    for i := 2; i <= n; i++ {
        factorization := factor.NewFactorization(i, max, primeIndex)
        factorization.Factor()
        fmt.Printf("TRACER i: %d %v \n", i, factorization)
    }
}

func process2(n int) {
    primeIndex := prime.New(n)

    a := 2
    b := 2
    max := n
    factorizationA := factor.NewFactorization(a, max, primeIndex)
    factorizationA.Factor()
    factorizationB := factor.NewFactorization(b, max, primeIndex)
    factorizationB.Factor()

    // test
    resultFactorization := factorizationA.Multiply(factorizationB)
    result := resultFactorization.GetPureFactorsString()
    fmt.Printf("TRACER result: %v\n", result)

    /*
    i := 3
    factorial := factorial.NewFactorial(i, max, primeIndex)
    factorial.Compute()
    fmt.Printf("TRACER factorial: %v \n", factorial.GetPureFactorsString())
    */
}

func process3(n int) {
    config := config.New(n)

    for i:= 2; i <= n; i++ {
        f := factorial.NewFactorial(i, config)
        f.Compute()
        fmt.Printf("TRACER f: %d -> %v \n", i, f.GetPureFactorsString())
    }
}

func main() {
    n := 10
    tmp1 := false
    if tmp1 { process1(n) }
    tmp2 := true
    if tmp2 { process2(n) }
    tmp3 := true
    if tmp3 { process3(n) }

    fmt.Println("Ready.")
}
