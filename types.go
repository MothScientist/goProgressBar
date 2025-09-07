package progressbar

// config structure for initial settings parameters
type config struct {
	spinner [10]string
	edges   [2]string
	colors  [2]string
}

// ProgressBar the interface defines methods for working with the progress bar
type ProgressBar interface {
    Update(percent uint8) (string)
    SetColors(newColors [2]string)
    SetSpinnerState(newState uint8) error
    SetPercent(newPercent uint8) error
    SetBarLen(newBarLen uint8)
    SetEdges(newEdges [2]string)
    SetCapString(newCapString string)
    SetEmptyString(newEmptyString string)
    WithPercent(show bool)
    WithSpinner(show bool)
    GetCurrentPercent() uint8
}

// progressBar structure containing all elements and settings of the progress bar
type progressBar struct {
	barLen       uint8
	percent      uint8
	color        [2]string
	edges        [2]string
	capString    string
	emptyString  string
	withPercent  bool
	withSpinner  bool
	spinnerState uint8
	spinnerLen   uint8
}