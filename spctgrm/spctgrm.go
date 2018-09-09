package spctgrm

import (
	"math"

	"github.com/mjibson/go-dsp/fft"
)

// Stft is a Short-Time Forier Transform.
// It can be used to return a complex Spectrogram of data.
func Stft(data []float64, windowSize, windowShift int) [][]complex128 {
	windowNum := len(data)/windowShift - (windowSize / windowShift) + 1
	freqArr := make([][]complex128, windowNum)
	for i := 0; i < windowNum; i++ {
		windowBegin := i * windowShift
		windowEnd := windowBegin + windowSize
		window := data[windowBegin:windowEnd]
		freqArr[i] = fft.FFTReal(window)
	}
	return flipAcrossHorizontal(transpose(freqArr))
}

func transpose(a [][]complex128) [][]complex128 {
	n := len(a)
	b := make([][]complex128, n)
	for i := 0; i < n; i++ {
		b[i] = make([]complex128, n)
		for j := 0; j < n; j++ {
			b[i][j] = a[j][i]
		}
	}
	return b
}

func flipAcrossHorizontal(a [][]complex128) [][]complex128 {
	n := len(a)
	for i := 0; i < n/2; i++ {
		j := len(a) - 1 - i
		a[i], a[j] = a[j], a[i]
	}
	return a
}

// NormSquared will converts complex numbers to the square of their norms
// on the asumption that real=x and imag=y
func NormSquared(graph [][]complex128) [][]float64 {
	outGraph := make([][]float64, len(graph))
	for i, row := range graph {
		outRow := make([]float64, len(row))
		for j, cell := range row {
			real := real(cell)
			imag := imag(cell)
			outRow[j] = 10 * math.Log10(real*real+imag*imag)
		}
		outGraph[i] = outRow
	}
	return outGraph
}
