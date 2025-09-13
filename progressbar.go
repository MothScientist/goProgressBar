package progressbar

import (
	"fmt"
	"strings"
)

const (
	ColorBlack         = "\033[30m"
	ColorRed           = "\033[31m"
	ColorGreen         = "\033[32m"
	ColorYellow        = "\033[33m"
	ColorBlue          = "\033[34m"
	ColorMagenta       = "\033[35m"
	ColorCyan          = "\033[36m"
	ColorWhite         = "\033[37m"
	ColorBrightBlack   = "\033[90m"
	ColorBrightRed     = "\033[91m"
	ColorBrightGreen   = "\033[92m"
	ColorBrightYellow  = "\033[93m"
	ColorBrightBlue    = "\033[94m"
	ColorBrightMagenta = "\033[95m"
	ColorBrightCyan    = "\033[96m"
	ColorBrightWhite   = "\033[97m"
	colorReset         = "\033[0m"
)

// baseConfig config containing settings for visual elements of the progress bar
var baseConfig = progressBarConfig{
	spinner:        spinners[0],
	edges:          [2]string{"[", "]"},
	colors:         [2]string{"", ColorGreen},
	fillers:        fillers[0],
	withPercent:    true,
	withSpinner:    true,
	reverseSpinner: false,
}

// GetNewProgressBar get the progress bar object
func GetNewProgressBar() *progressBar {
	return &progressBar{
		barLen:            50,
		percent:           0,
		spinnerState:      0,
		spinnerLen:        len(baseConfig.spinner),
		progressBarConfig: baseConfig,
	}
}

// progressBar returns the progress bar string with the current parameters
func (p *progressBar) render() string {
	occupancy := int((float32(p.barLen) / 100) * float32(p.percent))
	fill := p.barLen - occupancy

	// 0 1 2 3 4 5 6 ... bytes
	// ^-------^ ^
	// color     edge
	// if color != "", else:
	// 0 1 2 3 4 5 ... bytes
	// ^
	// edge

	startPart := fmt.Sprintf("%s%s", p.edges[0], p.colors[1])
	filledPart := p.fillers[0]
	emptyPart := p.fillers[1]
	endPart := fmt.Sprintf("%s%s%s", colorReset, p.colors[0], p.edges[1])

	var spinnerPart string
	if p.withSpinner {
		if !p.reverseSpinner {
			spinnerPart = fmt.Sprintf("%s ", p.spinner[p.spinnerState])

		} else {
			spinnerPart = fmt.Sprintf(" %s", p.spinner[p.spinnerState])
		}
	}

	var percentPart string
	if p.withPercent {
		if !p.reverseSpinner {
			percentPart = fmt.Sprintf(" %d%%", p.percent)
		} else {
			percentPart = fmt.Sprintf("%d%% ", p.percent)
		}
	}

	// Calculate how many bytes are required for the final string
	bytesToWrite := len(p.colors[0]) +
		len(startPart) + len(endPart) + // Progress bar edges
		len(spinnerPart) + len(percentPart) + // Optional parts
		(len(filledPart) * occupancy) + (len(emptyPart) * fill) + // Fillers
		len(colorReset) // Reset color
	var sb strings.Builder
	sb.Grow(bytesToWrite) // Allocate the required memory for the string

	sb.WriteString(p.colors[0])

	if !p.reverseSpinner && p.withSpinner {
		sb.WriteString(spinnerPart)
	} else if p.reverseSpinner && p.withPercent {
		sb.WriteString(percentPart)
	}

	sb.WriteString(startPart)
	sb.WriteString(strings.Repeat(filledPart, occupancy))
	sb.WriteString(strings.Repeat(emptyPart, fill))
	sb.WriteString(endPart)

	if !p.reverseSpinner && p.withPercent {
		sb.WriteString(percentPart)
	} else if p.reverseSpinner && p.withSpinner {
		sb.WriteString(spinnerPart)
	}

	sb.WriteString(colorReset)

	return sb.String()
}
