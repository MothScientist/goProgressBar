package progressbar

import "fmt"

// Update updates the percentage and returns the string for display
func (p *progressBar) Update(percent int) (string, error) {
	if percent < 0 || percent > 100 {
		return "", fmt.Errorf("the percent parameter must be no less than 0 and no more than 100, percent = %d", percent)
	}
	p.percent = percent
	if p.config.withSpinner {
		if percent == 0 {
			p.spinnerState = 0
		} else {
			p.spinnerState = (p.spinnerState + 1) % p.spinnerLen
		}
	}
	return p.render(), nil
}

// SetColors sets colors for frames, informers and fill
func (p *progressBar) SetColors(newColors [2]string) {
	p.config.colors = newColors
}

// SetPercent sets the fill percentage
func (p *progressBar) SetPercent(newPercent int) error {
	if newPercent > 100 {
		return fmt.Errorf("the percent variable must not be greater than 100, newPercent = %d", newPercent)
	}
	p.percent = newPercent
	return nil
}

// SetBarLen sets the length of the progress bar
func (p *progressBar) SetBarLen(newBarLen int) error {
	if newBarLen < 0 {
		return fmt.Errorf("newBarLen value must not be less than 0, newBarLen = %d", newBarLen)
	}
	p.barLen = newBarLen
	return nil
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
func (p *progressBar) SetSpinner(newSpinner []string) error {
	newSpinnerLen := len(newSpinner)
	if newSpinnerLen < 1 {
		return fmt.Errorf("spinner cut length cannot be less than 1, current length: %d", newSpinnerLen)
	}
	p.config.spinner = newSpinner
	p.spinnerLen = newSpinnerLen
	p.spinnerState = 0
	return nil
}

// WithPercent turns on/off the display of percentages
func (p *progressBar) WithPercent(show bool) {
	p.config.withPercent = show
}

// WithSpinner turns the spinner on/off
func (p *progressBar) WithSpinner(show bool) {
	p.config.withSpinner = show
}

// ReverseSpinner меняет спиннер и счетчик процентов местами
func (p *progressBar) ReverseSpinner(reverse bool) {
	p.config.reverseSpinner = reverse
}

// GetCurrentPercent get a set percentage
func (p *progressBar) GetCurrentPercent() int {
	return p.percent
}
