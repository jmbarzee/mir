package main

import (
	"fmt"
	"math"
	"os"

	"github.com/jmbarzee/mir/spctgrm"
	"github.com/mjibson/go-dsp/wav"
)

func main() {

	fName := "simple_loop"
	// fName := "simple_piano"

	floats, err := getFloatsFromWav(fmt.Sprintf("../audio/%s.wav", fName))
	if err != nil {
		panic(err)
	}

	// fName = "simple_sin"
	// floats = getFloatsFromSin(100000)

	err = dumpFloats1D(fmt.Sprintf("%s-audio", fName), floats)
	if err != nil {
		panic(err)
	}

	frameLen := 1024
	frameShift := frameLen / 4
	freqs := spctgrm.Stft(floats, frameLen, frameShift)
	freqsLog10 := spctgrm.NormSquared(freqs)

	err = dumpFloats2D(fmt.Sprintf("%s-amplitude", fName), freqsLog10)
	if err != nil {
		panic(err)
	}

}

func getFloatsFromSin(n int) []float64 {
	arr := make([]float64, n)
	for i := range arr {
		arr[i] = math.Sin(float64(i) / 10.0)
	}
	return arr
}

func getFloatsFromWav(fName string) ([]float64, error) {
	f, err := os.OpenFile(fName, os.O_RDONLY, 0666)
	if err != nil {
		return []float64{}, err
	}

	wav, err := wav.New(f)
	if err != nil {
		return []float64{}, err
	}

	allFloats := make([]float64, 0)
	n := 512 // TODO @jmbarzee consider changing n based on wav byte per sample
	newFloats, err := wav.ReadFloats(n)
	for err == nil {
		allFloats = append(allFloats, toFloat64(newFloats)...)
		newFloats, err = wav.ReadFloats(n)
	}

	return allFloats, nil
}

func toFloat64(floats32 []float32) []float64 {
	floats64 := make([]float64, len(floats32))
	for i, f := range floats32 {
		floats64[i] = float64(f)
	}
	return floats64
}
