package termutil

import (
	"image/color"
	"strings"
)

type CellAttributes struct {
	fgColour  color.Color
	bgColour  color.Color
	bold      bool
	dim       bool
	underline bool
	blink     bool
	inverse   bool
	hidden    bool
}

func (cellAttr *CellAttributes) reverseVideo() {
	oldFgColour := cellAttr.fgColour
	cellAttr.fgColour = cellAttr.bgColour
	cellAttr.bgColour = oldFgColour
}

// GetDiffANSI takes a previous cell attribute set and diffs to this one, producing the
// most efficient ANSI output to achieve the diff
func (cellAttr CellAttributes) GetDiffANSI(theme *Theme, prev CellAttributes) string {

	var segments []string

	// set fg
	if prev.fgColour != cellAttr.fgColour && cellAttr.fgColour != nil {
		segments = append(segments, theme.ColourToANSI(cellAttr.fgColour, false))
	} else if cellAttr.fgColour == nil {
		segments = append(segments, "\x1b[39m")
	}

	// set bg
	if prev.bgColour != cellAttr.bgColour && cellAttr.bgColour != nil {
		segments = append(segments, theme.ColourToANSI(cellAttr.bgColour, true))
	} else if cellAttr.bgColour == nil {
		segments = append(segments, "\x1b[49m")
	}

	// TODO add sequences for bold, dim, blink etc. diffs

	return strings.Join(segments, "")
}
