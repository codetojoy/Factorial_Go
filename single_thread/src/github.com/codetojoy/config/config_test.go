
package config

import (
    "testing"
    "github.com/codetojoy/prime"
)

func TestGetHighestPrime(t *testing.T) {
    cases := []struct {
        in int
        expected int
    }{
        { 2, 2 },
        { 3, 3 },
        { 4, 3 },
        { 5, 5 },
        { 6, 5 },
        { 7, 7 },
        { 8, 7 },
        { 9, 7 },
        { 10, 7 },
        { 40, 37 },
    }

    for _, c := range cases {
        primeIndex := prime.New(c.in)

        // test
        result := getHighestPrime(c.in, primeIndex)

        if result != c.expected {
            t.Errorf("getHighestPrime(%d) == %d, want %d", c.in, result, c.expected)
        }
    }
}
