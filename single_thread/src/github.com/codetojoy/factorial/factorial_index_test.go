
package factorial

import (
    "testing"

    "github.com/codetojoy/config"
)

func TestGetCompute(t *testing.T) {
    cases := []struct {
        in int
        expected string
    }{
        {2,"2: 1, "},
        {3,"2: 1, 3: 1, "},
        {4,"2: 3, 3: 1, "},
        {5,"2: 3, 3: 1, 5: 1, "},
    }

    max := 30
    config := config.New(max)
    factorialIndex := NewFactorialIndex(config)

    for _, c := range cases {
        // test
        factorial := factorialIndex.Get(c.in)

        result := factorial.GetPureFactorsString()

        if result != c.expected {
            t.Errorf("Get(%v) == %v, expected %v", c.in, result, c.expected)
        }
    }
}

func TestGetLookup(t *testing.T) {
    cases := []struct {
        in int
        expected string
    }{
        {2,"2: 1, "},
        {3,"2: 1, 3: 1, "},
        {4,"2: 3, 3: 1, "},
        {5,"2: 3, 3: 1, 5: 1, "},
    }

    max := 30
    config := config.New(max)
    factorialIndex := NewFactorialIndex(config)

    for _, c := range cases {
        factorial := factorialIndex.Get(c.in)

        // test
        factorial = factorialIndex.Get(c.in)

        result := factorial.GetPureFactorsString()

        if result != c.expected {
            t.Errorf("Get(%v) == %v, expected %v", c.in, result, c.expected)
        }
    }
}
