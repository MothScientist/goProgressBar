package progressbar

import "testing"

func TestBase(t *testing.T) {
	bar := GetNewProgressBar()
	bar.SetSpinner(Spinners[6])
	pg, _ := bar.Update(0)
	result := len(pg)
	expected := 172
	if result != expected {
		t.Errorf(`expected %d, want %d`, expected, result)
	}
}

func TestSpinnerState(t *testing.T) {
	bar := GetNewProgressBar()
	spinner := Spinners[0]
	bar.SetSpinner(spinner)
	spinnerLen := len(spinner)
	for i := range spinnerLen {
		pg, _ := bar.Update(i)
		result := string([]rune(pg)[0])
		expected := spinner[i]
		if result != spinner[i] {
			t.Errorf(`expected %s, want %s`, expected, result)
		}
	}
}

func TestChangeSpinner(t *testing.T) {
	bar := GetNewProgressBar()
	bar.SetSpinner(Spinners[6])
	// Check that baseConfig has not been touched
	if baseConfig.spinner[0] != Spinners[0][0] {
		t.Errorf(`baseConfig was affected by a change in the derived object, want %s`, baseConfig.spinner)
	}
}

func TestChangeFillers(t *testing.T) {}

func TestChangeColors(t *testing.T) {}

func TestChangeEdges(t *testing.T) {}

func TestChangeWithPercent(t *testing.T) {}

func TestChangeWithSpinner(t *testing.T) {}

func TestWithoutColors(t *testing.T) {
	bar := GetNewProgressBar()
	pg1, _ := bar.Update(0)
	result1 := len(pg1)
	bar.SetColors([2]string{"", ""})
	pg2, _ := bar.Update(0)
	result2 := len(pg2)
	diff := result1 - 5
	if diff != result2 {
		t.Errorf(`invalid number of bytes in string when color is missing %d, want %d`, result2, diff)
	}
}