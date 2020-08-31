
package worker

import (
    "github.com/codetojoy/factorial"
)

func SearchChunk(chunkLow int, chunkHigh int, factorialIndex *factorial.FactorialIndex, myChannel chan Payload) {
    var err error

    payloadResult := Payload{}

    for a := chunkLow; a <= chunkHigh && err == nil; a++ {
        factorialA := factorialIndex.Get(a)
        for b := 2; b <= a && err == nil; b++ {
            factorialB := factorialIndex.Get(b)
            for c := 2; c <= a && err == nil; c++ {
                factorialC := factorialIndex.Get(c)
                isMatch, innerErr := factorialA.IsProductMatch(factorialB, factorialC)

                if innerErr != nil {
                    err = innerErr
                } else if isMatch {
                    result := Result{A: a, B: b, C: c}
                    payloadResult.Results = append(payloadResult.Results, result)
                }
            }
        }
    }

    payloadResult.Err = err

    myChannel <- payloadResult
}
