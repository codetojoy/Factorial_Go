
package worker

import (
    "github.com/codetojoy/factorial"
)

type Payload struct {
    // in
    ChunkLow, ChunkHigh int
    FactorialIndex *factorial.FactorialIndex

    // out
    Results []Result
    Err error
}

type Result struct {
    A, B, C int
}
