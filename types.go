package progressbar

// config structure for initial settings parameters
type progressBarConfig struct {
	spinner        []string
	edges          [2]string
	colors         [2]string
	fillers        [2]string
	withPercent    bool
	withSpinner    bool
	reverseSpinner bool
}

// progressBar structure containing all elements and settings of the progress bar
type progressBar struct {
	barLen       int
	percent      int
	spinnerState int
	spinnerLen   int
	progressBarConfig
}
