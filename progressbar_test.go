package progressbar

import (
	"fmt"
	"testing"
)

func TestLen(t *testing.T) {
	bar := GetNewProgressBar()
	bar.SetSpinner(spinners[6])
	pg, _ := bar.Update(0)
	result := len(pg)
	expected := 172
	if result != expected {
		t.Errorf(`expected %d, want %d`, expected, result)
	}
}

func TestSpinnerState(t *testing.T) {
	bar := GetNewProgressBar()
	spinner := spinners[0]
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
	bar.SetSpinner(spinners[6])
	// Check that baseConfig has not been touched
	if baseConfig.spinner[0] != spinners[0][0] {
		t.Errorf(`baseConfig was affected by a change in the derived object, want %s`, baseConfig.spinner)
	}
}

func TestChangeFillers(t *testing.T) {}

func TestChangeColors(t *testing.T) {}

func TestChangeEdges(t *testing.T) {}

func TestChangeWithPercent0(t *testing.T) {
	bar := GetNewProgressBar()
	bar.SetColors([2]string{"", ""})
	result1, _ := bar.Update(0)
	bar.WithPercent(false)
	result2, _ := bar.Update(0)
	percentLen := 3
	if len(result2) != len(result1)-percentLen {
		t.Errorf(`некорректная длина прогресс бара без процентов: %d, want %d`, len(result2), len(result1)-percentLen)
	}
}

func TestChangeWithPercent50(t *testing.T) {
	bar := GetNewProgressBar()
	bar.SetColors([2]string{"", ""})
	result1, _ := bar.Update(50)
	bar.WithPercent(false)
	result2, _ := bar.Update(50)
	percentLen := 4
	if len(result2) != len(result1)-percentLen {
		t.Errorf(`некорректная длина прогресс бара без процентов: %d, want %d`, len(result2), len(result1)-percentLen)
	}
}

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

func TestZeroLenBar(t *testing.T) {
	bar := GetNewProgressBar()
	bar.SetColors([2]string{"", ""})
	err := bar.SetBarLen(0)
	if err != nil {
		t.Errorf(`error bar.SetBarLen(0): %v`, err)
	}
	pg1, _ := bar.Update(0)

	err = bar.SetBarLen(1)
	if err != nil {
		t.Errorf(`error bar.SetBarLen(1): %v`, err)
	}
	pg2, _ := bar.Update(0)

	bytesOneSymbol := 3
	if len(pg2) != len(pg1)+bytesOneSymbol {
		t.Errorf(`invalid string length %d, want %d`, len(pg2), len(pg1)+bytesOneSymbol)
	}
}

func TestReverseSpinner(t *testing.T) {
	bar := GetNewProgressBar()
	result1, _ := bar.Update(0)
	bar.ReverseSpinner(true)
	result2, _ := bar.Update(0)
	if len(result1) != len(result2) {
		t.Errorf(`некорректная длина прогресс бара %d, want %d`, len(result2), len(result1))
	}
}

func TestBase(t *testing.T) {
	bar := GetNewProgressBar()
	bar.SetSpinner(spinners[1])
	bar.SetEdges(edges[8])
	bar.SetFillers(fillers[3])
	for i := range 15 {
		result, _ := bar.Update(i + 50)
		fmt.Println(result)
	}
}
