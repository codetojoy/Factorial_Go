
package foo

import (
    "fmt"
)

type FooIndex struct {
    id int
}

func New(id int) FooIndex {
    return FooIndex{id: id}
}

func (f *FooIndex) String() string {
    return fmt.Sprintf("FooIndex id: %d", f.id)
}
