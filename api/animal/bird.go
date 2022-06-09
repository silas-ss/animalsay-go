package animal

import (
	"fmt"
	"strings"
)

type Bird struct {
	Name string
}

func (b Bird) Say(speech string) {
	b.printDialogBox(speech)
	b.printCow()
}

func (b Bird) printDialogBox(toSay string) {
	qtyCarac := len(toSay)

	border := strings.Repeat("-", qtyCarac)

	fmt.Println(border)
	fmt.Println(toSay)
	fmt.Println(border)
}

func (b Bird) printCow() {
	fmt.Println(`
    \
     \
	 __________
	 / ___  ___ \
	/ / @ \/ @ \ \
	\ \___/\___/ /\
	 \____\/____/||
	 /     /\\\\\//
	|     |\\\\\\
	 \      \\\\\\
	   \______/\\\\
		_||_||_
	`)
}
