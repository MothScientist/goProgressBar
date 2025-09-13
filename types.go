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

// ProgressBar the interface defines methods for working with the progress bar
type ProgressBar interface {
	Update(percent int) (string, error)
	SetColors(newColors [2]string)
	SetPercent(newPercent int) error
	SetBarLen(newBarLen int) error
	SetEdges(newEdges [2]string)
	SetFillers(newFillers [2]string)
	SetSpinner(newSpinner []string) error
	WithPercent(show bool)
	WithSpinner(show bool)
	ReverseSpinner(reverse bool)
	GetCurrentPercent() int
}

// progressBar structure containing all elements and settings of the progress bar
type progressBar struct {
	barLen       int
	percent      int
	config       progressBarConfig
	spinnerState int
	spinnerLen   int
}
