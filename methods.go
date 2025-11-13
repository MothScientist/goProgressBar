package progressbar

import "fmt"

// Update updates the percentage and returns the string for display
func (p *progressBar) Update(percent int) (string, error) {
	if percent < 0 || percent > 100 {
		return "", fmt.Errorf("the percent parameter must be no less than 0 and no more than 100, percent = %d", percent)
	}

	p.mu.Lock()
	p.percent = percent
	if p.withSpinner {
		if percent == 0 {
			p.spinnerState = 0
		} else {
			p.spinnerState = (p.spinnerState + 1) % p.spinnerLen
		}
	}
	p.mu.Unlock()

	return p.render(), nil
}

// SetColors sets colors for frames, informers and fill
func (p *progressBar) SetColors(newColors [2]string) {
	p.mu.Lock()
	p.colors = newColors
	p.mu.Unlock()
}

// SetPercent sets the fill percentage
func (p *progressBar) SetPercent(newPercent int) error {
	if newPercent > 100 {
		return fmt.Errorf("the percent variable must not be greater than 100, newPercent = %d", newPercent)
	}
	p.mu.Lock()
	p.percent = newPercent
	p.mu.Unlock()
	return nil
}

// SetBarLen sets the length of the progress bar
func (p *progressBar) SetBarLen(newBarLen int) error {
	if newBarLen < 0 {
		return fmt.Errorf("newBarLen value must not be less than 0, newBarLen = %d", newBarLen)
	}
	p.mu.Lock()
	p.barLen = newBarLen
	p.mu.Unlock()
	return nil
}

// SetEdges sets edge symbols
func (p *progressBar) SetEdges(newEdges [2]string) {
	p.mu.Lock()
	p.edges = newEdges
	p.mu.Unlock()
}

// SetFillers sets the fill and void characters
func (p *progressBar) SetFillers(newFillers [2]string) {
	p.mu.Lock()
	p.fillers = newFillers
	p.mu.Unlock()
}

// SetSpinner sets new spinner in progress bar
func (p *progressBar) SetSpinner(newSpinner []string) error {
	newSpinnerLen := len(newSpinner)
	if newSpinnerLen < 1 {
		return fmt.Errorf("spinner cut length cannot be less than 1, current length: %d", newSpinnerLen)
	}
	p.mu.Lock()
	p.spinner = newSpinner
	p.spinnerLen = newSpinnerLen
	p.spinnerState = 0
	p.mu.Unlock()
	return nil
}

// WithPercent turns on/off the display of percentages
func (p *progressBar) WithPercent(show bool) {
	p.mu.Lock()
	p.withPercent = show
	p.mu.Unlock()
}

// WithSpinner turns the spinner on/off
func (p *progressBar) WithSpinner(show bool) {
	p.mu.Lock()
	p.withSpinner = show
	p.mu.Unlock()
}

// ReverseSpinner меняет спиннер и счетчик процентов местами
func (p *progressBar) ReverseSpinner(reverse bool) {
	p.mu.Lock()
	p.reverseSpinner = reverse
	p.mu.Unlock()
}

// GetCurrentPercent get a set percentage
func (p *progressBar) GetCurrentPercent() (percent int) {
	p.mu.RLock()
	percent = p.percent
	p.mu.RUnlock()
	return
}
