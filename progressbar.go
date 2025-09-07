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

var defaultConfig = config{
	spinner: [10]string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"},
	edges:   [2]string{"[", "]"},
	colors:  [2]string{ColorWhite, ColorGreen},
}

// GetNewProgressBar get the progress bar object
func GetNewProgressBar() ProgressBar {
	return &progressBar{
		barLen:       50,
		percent:      0,
		color:        defaultConfig.colors,
		edges:        defaultConfig.edges,
		capString:    "█",
		emptyString:  "░",
		withPercent:  true,
		withSpinner:  true,
		spinnerState: 0,
		spinnerLen:   uint8(len(defaultConfig.spinner)),
	}
}

// progressBar returns the progress bar string with the current parameters
func (p *progressBar) render() string {
	occupancy := uint8((float32(p.barLen) / 100) * float32(p.percent))
	fill := p.barLen - occupancy
	startPart := fmt.Sprintf("%s%s%s", p.color[0], p.edges[0], p.color[1])
	filledPart := p.capString
	emptyPart := p.emptyString
	endPart := fmt.Sprintf("%s%s%s%s", colorReset, p.color[0], p.edges[1], colorReset)

	var sb strings.Builder

	if p.withSpinner {
		sb.WriteString(fmt.Sprintf("%s%s%s ", p.color[0], defaultConfig.spinner[p.spinnerState], colorReset))
	}

	sb.WriteString(startPart)
	sb.WriteString(strings.Repeat(filledPart, int(occupancy)))
	sb.WriteString(strings.Repeat(emptyPart, int(fill)))
	sb.WriteString(endPart)

	if p.withPercent {
		sb.WriteString(fmt.Sprintf(" %s%d%%%s", p.color[0], p.percent, colorReset))
	}

	return sb.String()
}
