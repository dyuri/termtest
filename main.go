package main

import (
	"fmt"
	"time"

	"github.com/dyuri/termtest/termutil"
)

func render(b *termutil.Buffer) {
	var lastCellAttr termutil.CellAttributes

	for line := range b.ViewHeight() {
		for row := range b.ViewWidth() {
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
	terminal := termutil.New(termutil.WithCommand("exa", "-la", "/home/dyuri/alma"))
	updateChan := make(chan struct{})
	// closeChan := make(chan struct{})

	terminal.Run(updateChan, 20, 80)

	go func() {
		for {
			<-time.After(1 * time.Second)
			updateChan <- struct{}{}
		}
	}()

	<-time.After(5 * time.Second)

	render(terminal.GetActiveBuffer())
	fmt.Printf("Active buffer length: %d\n", terminal.GetActiveBuffer().ViewHeight())
	fmt.Printf("Active buffer scroll offset: %d\n", terminal.GetActiveBuffer().GetScrollOffset())
}
