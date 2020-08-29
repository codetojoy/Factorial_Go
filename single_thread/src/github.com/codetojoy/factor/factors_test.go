
package factor

import (
    "testing"

    "github.com/codetojoy/config"
)

func TestPut(t *testing.T) {
    config := config.New(10)
    factors := NewFactors(config)

    // test
    factors.Put(2,2);
    factors.Put(3,3);

    result := factors.String()
    expected := "2: 2, 3: 3, "

    if result != expected {
        t.Errorf("Factors.Put() == %v, want %v", result, expected)
    }
}

func TestEquals(t *testing.T) {
    cases := []struct {
        inA []int
        inB []int
        expected bool
    }{
        { []int{1,1}, []int{1,1}, true },
        { []int{1,1,0}, []int{1,1}, false },
        { []int{3,4,5}, []int{3,4,5}, true },
        { []int{3,3,3}, []int{3,4,5}, false },
    }

    config := config.New(5)

    for _, c := range cases {
        factorsA := NewFactorsForTesting(config, c.inA)
        factorsB := NewFactorsForTesting(config, c.inB)

        // test
        result := factorsA.Equals(factorsB)

        if result != c.expected {
            t.Errorf("[%v].Equals(%v) == %v, expected %v", c.inA, c.inB, result, c.expected)
        }
    }
}

func TestMultiply(t *testing.T) {
    cases := []struct {
        inA []int
        inB []int
        expected []int
    }{
        { []int{1,0,0,0,0}, []int{0,1,0,0,0}, []int{1,1,0,0,0} },
        { []int{1,1,0,0,0}, []int{1,1,0,0,1}, []int{2,2,0,0,1} },
    }

    config := config.New(5)

    for _, c := range cases {
        factorsA := NewFactorsForTesting(config, c.inA)
        factorsB := NewFactorsForTesting(config, c.inB)
        factorsExpected := NewFactorsForTesting(config, c.expected)

        // test
        result, e := factorsA.Multiply(factorsB)

        if e == nil && result.Equals(factorsExpected) {
            t.Errorf("[%v].Multiply(%v) == %v, expected %v", c.inA, c.inB, result, c.expected)
        }
    }
}

func TestMultiplyError(t *testing.T) {
    configA := config.New(5)
    configB := config.New(3)
    factorsA := NewFactorsForTesting(configA, []int{1,0,0,0,0})
    factorsB := NewFactorsForTesting(configB, []int{1,0,0})

    // test
    _, e := factorsA.Multiply(factorsB)

    if e == nil {
        t.Errorf("Factors.Multiply() expected error")
    }
}

func TestStringError(t *testing.T) {
    config := config.New(5)
    factors := NewFactorsForTesting(config, []int{})

    // test
    result := factors.String()

    expected := "ERROR: NONE"

    if result != expected {
        t.Errorf("Factors.Multiply() expected %v", expected)
    }
}

func TestIsProductMatch(t *testing.T) {
    cases := []struct {
        inA []int
        inB []int
        inC []int
        expected bool
    }{
        { []int{1,0,0,0,1}, []int{1,0,0,0,0}, []int{0,0,0,0,1}, true },
        { []int{4,3,2,1,0}, []int{2,3,1,0,0}, []int{2,0,1,1,0}, true },
    }

    config := config.New(5)

    for _, c := range cases {
        factorsA := NewFactorsForTesting(config, c.inA)
        factorsB := NewFactorsForTesting(config, c.inB)
        factorsC := NewFactorsForTesting(config, c.inC)

        // test
        result, e := factorsA.IsProductMatch(factorsB, factorsC)

        if e != nil || (! result) {
            t.Errorf("[%v].IsProductMatch(%v) == %v, expected %v", c.inA, c.inB, c.inC, c.expected)
        }
    }
}

func TestIsProductMatchError(t *testing.T) {
    configAC := config.New(5)
    configB := config.New(3)

    factorsA := NewFactorsForTesting(configAC, []int{1,0,0,0,0})
    factorsB := NewFactorsForTesting(configB, []int{1,0,0})
    factorsC := NewFactorsForTesting(configAC, []int{1,0,0,0,0})

    // test
    _, e := factorsA.IsProductMatch(factorsB, factorsC)

    if e == nil {
        t.Errorf("Factors.Multiply() expected error")
    }
}
