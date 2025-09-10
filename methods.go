package progressbar

import "fmt"

// Update updates the percentage and returns the string for display
func (p *progressBar) Update(percent uint8) string {
	p.percent = percent
	if p.config.withSpinner {
		if percent == 0 {
			p.spinnerState = 0
		} else {
			p.spinnerState = (p.spinnerState + 1) % p.spinnerLen
		}
	}
	return p.render()
}


// SetColors sets colors for frames, informers and fill
func (p *progressBar) SetColors(newColors [2]string) {
	p.config.colors = newColors
}

// SetPercent sets the fill percentage
func (p *progressBar) SetPercent(newPercent uint8) error {
	if newPercent > 100 {
		return fmt.Errorf("переменная percent не должна быть больше 100")
	}
	p.percent = newPercent
	return nil
}

// SetBarLen sets the length of the progress bar
func (p *progressBar) SetBarLen(newBarLen uint8) {
	p.barLen = newBarLen
}

// SetEdges sets edge symbols
func (p *progressBar) SetEdges(newEdges [2]string) {
	p.config.edges = newEdges
}

// SetFillers sets the fill and void characters
func (p *progressBar) SetFillers(newFillers [2]string) {
	p.config.fillers = newFillers
}

// SetSpinner sets new spinner in progress bar
func (p *progressBar) SetSpinner(newSpinner []string) {
	p.config.spinner = newSpinner
	p.spinnerLen = uint8(len(newSpinner))
	p.spinnerState = 0
}

// WithPercent turns on/off the display of percentages
func (p *progressBar) WithPercent(show bool) {
	p.config.withPercent = show
}

// WithSpinner turns the spinner on/off
func (p *progressBar) WithSpinner(show bool) {
	p.config.withSpinner = show
}

// GetCurrentPercent get a set percentage
func (p *progressBar) GetCurrentPercent() uint8 {
	return p.percent
}