package progressbar

import (
	"fmt"
	"strings"
)

const (
	ColorWhite  = "\033[37m"
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorBlue   = "\033[34m"
	ColorPurple = "\033[35m"
	ColorCyan   = "\033[36m"
	ColorGray   = "\033[37m"
	colorReset  = "\033[0m"
)

// Spinners built-in spinners for progress bar
var Spinners = map[uint8][]string{
	0: {"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"},
	1: {"⣾", "⣽", "⣻", "⢿", "⡿", "⣟", "⣯", "⣷"},
	2: {"▁", "▂", "▃", "▄", "▅", "▆", "▇"},
	3: {"*--", "-*-", "--*"},
	4: {".  ", ".. ", "...", " ..", "  ."},
	5: {"*  ", "** ", "***", " **", "  *"},
	6: {"*  ", " * ", "  *"},
}

// Fillers built-in fillers for progress bar
var Fillers = map[uint8][2]string{
	0: {"█", "░"},
	1: {"|", " "},
	2: {"#", " "},
}

// baseConfig config containing settings for visual elements of the progress bar
var baseConfig = progressBarConfig{
	spinner:     Spinners[0],
	edges:       [2]string{"[", "]"},
	colors:      [2]string{"", ColorGreen},
	fillers:     Fillers[0],
	withPercent: true,
	withSpinner: true,
}

// GetNewProgressBar get the progress bar object
func GetNewProgressBar() ProgressBar {
	return &progressBar{
		barLen:       50,
		percent:      0,
		config:       baseConfig,
		spinnerState: 0,
		spinnerLen:   len(baseConfig.spinner),
	}
}

// progressBar returns the progress bar string with the current parameters
func (p *progressBar) render() string {
	occupancy := int((float32(p.barLen) / 100) * float32(p.percent))
	fill := p.barLen - occupancy

	startPart := fmt.Sprintf("%s%s%s", p.config.colors[0], p.config.edges[0], p.config.colors[1])
	filledPart := p.config.fillers[0]
	emptyPart := p.config.fillers[1]
	endPart := fmt.Sprintf("%s%s%s", colorReset, p.config.colors[0], p.config.edges[1])

	var spinnerPart string
	if p.config.withSpinner {
		spinnerPart = fmt.Sprintf("%s%s ", p.config.colors[0], p.config.spinner[p.spinnerState])
		// 0 1 2 3 4 5 6 ... bytes
		// ^-------^ ^
		// color     spinner
		// if color != "", else:
		// 0 1 2 3 4 5 ... bytes
		// ^
		// spinner
	}

	var percentPart string
	if p.config.withPercent {
		percentPart = fmt.Sprintf(" %d%%", p.percent)
	}

	// Calculate how many bytes are required for the final string
	bytesToWrite := len(startPart) + len(endPart) + // Progress bar edges
		len(spinnerPart) + len(percentPart) + // Optional parts
		(len(filledPart) * occupancy) + (len(emptyPart) * fill) + // Fillers
		len(colorReset) // Reset color
	var sb strings.Builder
	sb.Grow(bytesToWrite) // Allocate the required memory for the string

	if p.config.withSpinner {
		sb.WriteString(spinnerPart)
	}

	sb.WriteString(startPart)
	sb.WriteString(strings.Repeat(filledPart, occupancy))
	sb.WriteString(strings.Repeat(emptyPart, fill))
	sb.WriteString(endPart)

	if p.config.withPercent {
		sb.WriteString(percentPart)
	}

	sb.WriteString(colorReset)

	return sb.String()
}
