package sequences

import (
	"testing"
)

func TestCalculateAP(t *testing.T) {
	ap, err := CalcAP(2, 4, 6)
	if err != nil || ap != 8 {
		t.Fail()
	}
}

func TestCalculateGP(t *testing.T) {
	val, err := CalcGP(2, 4, 8)
	if err != nil || val != 16 {
		t.Fail()
	}
}

func TestCalculateAPSequence(t *testing.T) {
	val, err := CalcAPs(2, 4, 6, 5)
	if err != nil {
		t.Fail()
	}
	expected := []float64{8, 10, 12, 14, 16}
	for i := range val {
		if val[i] != expected[i] {
			t.Fail()
		}
	}
}

func TestCalculateGPSequence(t *testing.T) {
	val, err := CalcGPs(2, 4, 8, 5)
	if err != nil {
		t.Fail()
	}
	expected := []float64{16, 32, 64, 128, 256}
	for i := range val {
		if val[i] != expected[i] {
			t.Fail()
		}
	}
}

func TestCalculateSplitAPSequences(t *testing.T) {
	val, err := CalcSplitAPs([]float64{2, 4, 6, 4, 6, 8, 6, 8, 10, 8, 10, 12})
	if err != nil {
		t.Fail()
	}
	expected := []float64{2, 4, 6, 8,  4, 6, 8, 10,  6, 8, 10, 12,  8, 10, 12, 14}
	for i := range val {
		if val[i] != expected[i] {
			t.Fail()
		}
	}
}

func TestCalculateSplitGPSequences(t *testing.T) {
	val, err := CalcSplitGPs([]float64{2, 4, 8, 10, 30, 90, 100, 500, 2500})
	if err != nil {
		t.Fail()
	}
	expected := []float64{2, 4, 8, 16, 10, 30, 90, 270, 100, 500, 2500, 12500}
	for i := range val {
		if val[i] != expected[i] {
			t.Fail()
		}
	}
}