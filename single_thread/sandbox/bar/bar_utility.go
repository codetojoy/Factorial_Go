
package bar

import (
    "fmt"
)

type BarUtility struct {
    id int
}

func New(id int) BarUtility {
    return BarUtility{id: id}
}

func (bu *BarUtility) String() string {
    return fmt.Sprintf("BarUtility id: %d", bu.id)
}
