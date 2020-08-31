
package factorial

import (
    "testing"

    "github.com/codetojoy/config"
)

func TestCompute(t *testing.T) {
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

    for _, c := range cases {
        factorial := NewFactorial(c.in, config)

        // test
        factorial.Compute()
        result := factorial.GetPureFactorsString()

        if result != c.expected {
            t.Errorf("Compute(%v) == %v, expected %v", c.in, result, c.expected)
        }
    }
}

func TestString(t *testing.T) {
    max := 30
    config := config.New(max)
    n := 5
    factorial := NewFactorial(n, config)
    factorial.Compute()

    // test
    result := factorial.String()

    expected := "5! => -1 :: 2: 3, 3: 1, 5: 1, "

    if result != expected {
        t.Errorf("String() expected: %v got: %v", expected, result)
    }
}
