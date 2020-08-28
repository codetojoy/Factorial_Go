
package main

import (
    "fmt"
    "./factor"
    "./factorial"
    "./prime"
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

    i := 3
    factorial := factorial.NewFactorial(i, max, primeIndex)
    factorial.Compute()
    fmt.Printf("TRACER factorial: %v \n", factorial.GetPureFactorsString())
    /*
    max := 10
    i := 3
    factorial := NewFactorial(i, max, primeIndex)
    factorial.Compute()
    fmt.Printf("TRACER factorial: %v \n", factorial.GetPureFactorsString())
    x := 5
    ceiling := (x/2) + 1
    f := NewFactorization(x, primeIndex)
    f.Factor()
    fmt.Printf("TRACER x: %d ceiling: %d f: %v \n", x, ceiling, f.GetPureFactorsString())
    for i := 2; i <= 3; i++ {
    }
    */
}

func main() {
    n := 20
    tmp1 := false
    if tmp1 { process1(n) }
    tmp2 := true
    if tmp2 { process2(n) }
    /*
    fmt.Println("TRACER ------------------")
    primeIndex := PrimeIndex{}
    primeIndex.Initialize(20);
    fmt.Printf("TRACER p[%d]: %d\n", 0, primeIndex.GetPrimeForIndex(0))
    fmt.Printf("TRACER p[%d]: %d\n", 1, primeIndex.GetPrimeForIndex(1))
    for i := 0; i < 20; i++ {
        fmt.Printf("TRACER p[%d]: %d\n", i, primeIndex.GetPrimeForIndex(i))
    }
    */
    fmt.Println("Ready.")
}
