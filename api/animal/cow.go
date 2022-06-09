package animal

import (
	"fmt"
	"strings"
)

type Cow struct {
	Name string
}

func (c Cow) Say(speech string) {
	c.printDialogBox(speech)
	c.printCow()
}

func (c Cow) printDialogBox(toSay string) {
	qtyCarac := len(toSay)

	border := strings.Repeat("-", qtyCarac)

	fmt.Println(border)
	fmt.Println(toSay)
	fmt.Println(border)
}

func (c Cow) printCow() {
	fmt.Println(`
    \
     \
      ------
     / 0 0 /
    / ._. / \____________
    ------\              \
          | _  ______     \
          || ||      ||w|| \
          || ||      || ||
	`)
}
