
package main

import (
    "fmt"
)

type MyIndex struct {
    id int
}

func New(id int) MyIndex {
    return MyIndex{id: id}
}

func (mi *MyIndex) String() string {
    return fmt.Sprintf("MyIndex id: %d", mi.id)
}
