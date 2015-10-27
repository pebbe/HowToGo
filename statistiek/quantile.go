package main

import (
	"fmt"
)

func main() {
	a := []float64{1, 2, 3, 4, 5}
	fmt.Printf("Want: 2.00  3.00  4.00\n")
	fmt.Printf("Got:  %.2f  %.2f  %.2f\n\n", quant(a, .25), quant(a, .5), quant(a, .75))
	a = []float64{1, 2, 3, 4, 5, 6}
	fmt.Printf("Want: 2.25  3.50  4.75\n")
	fmt.Printf("Got:  %.2f  %.2f  %.2f\n\n", quant(a, .25), quant(a, .5), quant(a, .75))
	a = []float64{1, 2, 3, 4, 5, 6, 7}
	fmt.Printf("Want: 2.50  4.00  5.50\n")
	fmt.Printf("Got:  %.2f  %.2f  %.2f\n\n", quant(a, .25), quant(a, .5), quant(a, .75))
	a = []float64{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Printf("Want: 2.75  4.50  6.25\n")
	fmt.Printf("Got:  %.2f  %.2f  %.2f\n\n", quant(a, .25), quant(a, .5), quant(a, .75))
	a = []float64{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("Want: 3.00  5.00  7.00\n")
	fmt.Printf("Got:  %.2f  %.2f  %.2f\n\n", quant(a, .25), quant(a, .5), quant(a, .75))

	a = []float64{1, 2, 4, 7, 11}
	fmt.Printf("Want: 2.00  4.00  7.00\n")
	fmt.Printf("Got:  %.2f  %.2f  %.2f\n\n", quant(a, .25), quant(a, .5), quant(a, .75))
	a = []float64{1, 2, 4, 7, 11, 16}
	fmt.Printf("Want: 2.50  5.50  10.00\n")
	fmt.Printf("Got:  %.2f  %.2f  %.2f\n\n", quant(a, .25), quant(a, .5), quant(a, .75))
	a = []float64{1, 2, 4, 7, 11, 16, 22}
	fmt.Printf("Want: 3.00  7.00  13.50\n")
	fmt.Printf("Got:  %.2f  %.2f  %.2f\n\n", quant(a, .25), quant(a, .5), quant(a, .75))
	a = []float64{1, 2, 4, 7, 11, 16, 22, 29}
	fmt.Printf("Want: 3.50  9.00  17.50\n")
	fmt.Printf("Got:  %.2f  %.2f  %.2f\n\n", quant(a, .25), quant(a, .5), quant(a, .75))
	a = []float64{1, 2, 4, 7, 11, 16, 22, 29, 36}
	fmt.Printf("Want: 4.00  11.00  22.00\n")
	fmt.Printf("Got:  %.2f  %.2f  %.2f\n", quant(a, .25), quant(a, .5), quant(a, .75))
}

func quant(f []float64, part float64) (result float64) {
	left := float64(len(f)-1) * part
	i := int(left)
	fr := left - float64(i)
	return f[i]*(1-fr) + f[i+1]*fr
}
