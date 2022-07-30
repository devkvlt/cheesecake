package main

import (
	"fmt"

	cc "github.com/devkvlt/cheesecake/cheesecake"
)

func main() {

	// Ruy Lopez
	b := cc.MakeDefaultBoard()
	b.Move(cc.E2, cc.E4)
	b.Move(cc.E7, cc.E5)
	b.Move(cc.G1, cc.F3)
	b.Move(cc.B8, cc.C6)
	b.Move(cc.F1, cc.B5)

	fmt.Println(b)

}
