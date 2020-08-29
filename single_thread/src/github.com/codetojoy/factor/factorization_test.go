
package factor

import (
    "testing"

    "github.com/codetojoy/config"
)

func TestFactorizationMultiply(t *testing.T) {
    cases := []struct {
        inA int
        inB int
        expected string
    }{
        {2,2,"2: 2, "},
        {2,3,"2: 1, 3: 1, "},
        {8,6,"2: 4, 3: 1, "},
    }

    max := 30
    config := config.New(max)

    for _, c := range cases {
        factorizationA := NewFactorization(c.inA, config)
        factorizationA.Factor()
        factorizationB := NewFactorization(c.inB, config)
        factorizationB.Factor()

        // test
        resultFactorization := factorizationA.Multiply(factorizationB)
        result := resultFactorization.GetPureFactorsString()

        if result != c.expected {
            t.Errorf("[%v].Multiply(%v) == %v, expected %v", c.inA, c.inB, result, c.expected)
        }
    }
}

func TestFactor(t *testing.T) {
    cases := []struct {
        in int
        expected string
    }{
        {2,"2 :: 2: 1, "},
        {3,"3 :: 3: 1, "},
        {4,"4 :: 2: 2, "},
        {5,"5 :: 5: 1, "},
        {6,"6 :: 2: 1, 3: 1, "},
        {7,"7 :: 7: 1, "},
        {8,"8 :: 2: 3, "},
        {9,"9 :: 3: 2, "},
        {10,"10 :: 2: 1, 5: 1, "},
        {28,"28 :: 2: 2, 7: 1, "},
    }

    max := 30
    config := config.New(max)

    for _, c := range cases {
        factorization := NewFactorization(c.in, config)

        // test
        factorization.Factor()
        result := factorization.String()

        if result != c.expected {
            t.Errorf("Factor(%v) == %v, expected %v", c.in, result, c.expected)
        }
    }
}
