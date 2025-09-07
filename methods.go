package progressbar

import "fmt"

// Update updates the percentage and returns the string for display
func (p *progressBar) Update(percent uint8) string {
	p.percent = percent
	if p.withSpinner {
		p.spinnerState = (p.spinnerState + 1) % p.spinnerLen
	}
	return p.render()
}


// SetColors sets colors for frames, informers and fill
func (p *progressBar) SetColors(newColors [2]string) {
	p.color = newColors
}

// SetSpinnerState sets the spinner state
func (p *progressBar) SetSpinnerState(newState uint8) error {
	if p.spinnerState >= p.spinnerLen {
		return fmt.Errorf("параметр spinner не должен быть более %d, сейчас он равен %d", p.spinnerLen, p.spinnerState)
	}
	p.spinnerState = newState
	return nil
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
	p.edges = newEdges
}

// SetCapString sets the fill character
func (p *progressBar) SetCapString(newCapString string) {
	p.capString = newCapString
}

// SetEmptyString sets the void symbol
func (p *progressBar) SetEmptyString(newEmptyString string) {
	p.emptyString = newEmptyString
}

// WithPercent turns on/off the display of percentages
func (p *progressBar) WithPercent(show bool) {
	p.withPercent = show
}

// WithSpinner turns the spinner on/off
func (p *progressBar) WithSpinner(show bool) {
	p.withSpinner = show
}

// GetCurrentPercent get a set percentage
func (p *progressBar) GetCurrentPercent() uint8 {
	return p.percent
}