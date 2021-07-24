package sequences

import (
	"errors"
	"math"
)

func calcNextSequence(a, b, c float64, n int, apSeq bool) (float64, error) {
	if apSeq {
		if c-b != b-a {
			return 0, errors.New("not a valid sequence")
		}
		r := c - b
		return a + (float64(n)-1)*r, nil
	} else {
		if c/b != b/a {
			return 0, errors.New("not a valid sequence")
		}
		r := c / b
		return a * math.Pow(r, float64(n)-1), nil
	}
}

func calcNextSequences(a, b, c float64, n int, apSeq bool) ([]float64, error) {
	seq := make([]float64, n)
	for i := 0; i < n; i++ {
		val, err := calcNextSequence(a, b, c, i+4, apSeq)
		if err != nil {
			return nil, err
		} else {
			seq[i] = val
		}
	}
	return seq, nil
}

func calcSplitSequences(arr []float64, apSeq bool) ([]float64, error) {
	var seq []float64
	for i := 0; i < len(arr); i += 3 {
		a, b, c := arr[i], arr[i+1], arr[i+2]
		d, err := calcNextSequence(a, b, c, 4, apSeq)
		if err != nil {
			return nil, err
		} else {
			seq = append(seq, a, b, c, d)
		}
	}
	return seq, nil
}

func CalcAP(a, b, c float64) (float64, error) {
	return calcNextSequence(a, b, c, 4, true)
}

func CalcGP(a, b, c float64) (float64, error) {
	return calcNextSequence(a, b, c, 4, false)
}

func CalcAPs(a, b, c float64, n int) ([]float64, error) {
	return calcNextSequences(a, b, c, n, true)
}

func CalcGPs(a, b, c float64, n int) ([]float64, error) {
	return calcNextSequences(a, b, c, n, false)
}

func CalcSplitAPs(arr []float64) ([]float64, error) {
	return calcSplitSequences(arr, true)
}

func CalcSplitGPs(arr []float64) ([]float64, error) {
	return calcSplitSequences(arr, false)
}
