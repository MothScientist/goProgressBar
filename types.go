package progressbar

// config structure for initial settings parameters
type progressBarConfig struct {
	spinner     []string
	edges       [2]string
	colors      [2]string
	fillers     [2]string
	withPercent bool
	withSpinner bool
}

// ProgressBar the interface defines methods for working with the progress bar
type ProgressBar interface {
    Update(percent uint8) string
    SetColors(newColors [2]string)
    SetPercent(newPercent uint8) error
    SetBarLen(newBarLen uint8)
    SetEdges(newEdges [2]string)
    SetFillers(newFillers [2]string)
	SetSpinner(newSpinner []string)
    WithPercent(show bool)
    WithSpinner(show bool)
    GetCurrentPercent() uint8
}

// progressBar structure containing all elements and settings of the progress bar
type progressBar struct {
	barLen       uint8
	percent      uint8
	config       progressBarConfig
	spinnerState uint8
	spinnerLen   uint8
}
