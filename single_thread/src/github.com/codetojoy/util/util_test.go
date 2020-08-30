
package util

import (
    "testing"
)

func TestGetChunk(t *testing.T) {
    cases := []struct {
        index, expectedLow, expectedHigh int
    }{
        { 1, 2, 20 },
        { 2, 20, 40 },
        { 3, 40, 60 },
        { 4, 60, 70 },
    }

    const max = 70
    const chunkSize = 20

    for _, c := range cases {
        // test
        resultLow, resultHigh := GetChunk(c.index, chunkSize, max)

        if resultLow != c.expectedLow || resultHigh != c.expectedHigh {
            t.Errorf("getChunk() == (%d,%d), want (%d,%d)", resultLow, resultHigh,
                c.expectedLow, c.expectedHigh)
        }
    }
}
