package progressbar

import (
	"math"
	"testing"
)

func TestEdges(t *testing.T) {
	originalLen := len(edges)
	expEdges := GetEdges()
	expEdgesLen := len(expEdges)
	if originalLen != expEdgesLen {
		t.Errorf("incorrect map length %d, want %d", expEdgesLen, originalLen)
	}
}

func TestFillers(t *testing.T) {
	originalLen := len(fillers)
	expFillers := GetFillers()
	expFillersLen := len(expFillers)
	if originalLen != expFillersLen {
		t.Errorf("incorrect map length %d, want %d", expFillersLen, originalLen)
	}
}

func TestSpinners(t *testing.T) {
	originalLen := len(spinners)
	expSpinners := GetSpinners()
	expSpinnersLen := len(expSpinners)
	if originalLen != expSpinnersLen {
		t.Errorf("incorrect map length %d, want %d", expSpinnersLen, originalLen)
	}
}

func TestChangeSpinnersMap(t *testing.T) {
	originalLen := len(spinners)

	expSpinners := GetSpinners()
	expSpinners[math.MaxInt32] = []string{"1", "2", "3"}
	expSpinnersLen := len(expSpinners)

	if originalLen == expSpinnersLen {
		t.Errorf("incorrect map length %d, want %d", expSpinnersLen, originalLen - 1)
	}
}