package main

import (
	"fmt"
	"math"
	"os"
)

func dumpComplex2D(fName string, graph [][]complex128) error {
	realf, err := os.OpenFile("real-"+fName, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	imagf, err := os.OpenFile("imag-"+fName, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}

	for _, row := range graph {
		for _, cell := range row {
			reals := fmt.Sprintf("%.3g", math.Log10(real(cell)))
			fmt.Fprintf(realf, "%s, ", reals)
			imags := fmt.Sprintf("%.3g", math.Log10(imag(cell)))
			fmt.Fprintf(imagf, "%s, ", imags)
		}
		fmt.Fprintf(realf, "\n")
		fmt.Fprintf(imagf, "\n")
	}
	return nil
}

func dumpFloats2D(fName string, graph [][]float64) error {
	fName = fmt.Sprintf("../dump/%s.csv", fName)
	f, err := os.OpenFile(fName, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}

	for _, row := range graph {
		for _, cell := range row {
			s := fmt.Sprintf("%.3g", cell)
			// s = strings.TrimPrefix(s, "(")
			// s = strings.TrimSuffix(s, ")")
			fmt.Fprintf(f, "%s, ", s)
		}
		fmt.Fprintf(f, "\n")
	}
	return nil
}

func dumpFloats1D(fName string, arr []float64) error {
	fName = fmt.Sprintf("../dump/%s.csv", fName)
	f, err := os.OpenFile(fName, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}

	for _, cell := range arr {
		fmt.Fprintf(f, "%.3g, ", cell)
	}
	return nil
}
