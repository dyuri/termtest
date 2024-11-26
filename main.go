package main

import (
	"fmt"
	"time"

	"github.com/dyuri/termtest/termutil"
)

func render(b *termutil.Buffer) {
	var lastCellAttr termutil.CellAttributes

	for line := range uint16(b.Height()) {
		for row := range b.Width() {
			cell := b.GetCell(row, line)
			if cell != nil {
				r := cell.Rune().Rune
				if r < 0x20 {
					r = '.'
				}
				sgr := cell.Attr().GetDiffANSI(&termutil.DefaultTheme, lastCellAttr)
				fmt.Printf("%s", sgr + string(r))
				lastCellAttr = cell.Attr()
			} else {
				cellAttr := termutil.CellAttributes{}
				sgr := cellAttr.GetDiffANSI(&termutil.DefaultTheme, lastCellAttr)
				fmt.Printf("%s•", sgr)
				lastCellAttr = cellAttr
			}
		}
		fmt.Println()
	}
}

func main() {
	terminal := termutil.New(termutil.WithCommand("/bin/ls", "--color"))
	updateChan := make(chan struct{})
	// closeChan := make(chan struct{})

	terminal.Run(updateChan, 10, 80)

	go func() {
		updateChan <- struct{}{}
	}()

	<-time.After(1 * time.Second)

	render(terminal.GetActiveBuffer())
}
