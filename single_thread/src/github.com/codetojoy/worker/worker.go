
package worker

func SearchChunk(payload Payload) Payload {
    var err error

    chunkLow := payload.ChunkLow
    chunkHigh := payload.ChunkHigh
    factorialIndex := payload.FactorialIndex

    payloadResult := Payload{ChunkLow: chunkLow, ChunkHigh: chunkHigh, FactorialIndex: factorialIndex}

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

    return payloadResult
}
