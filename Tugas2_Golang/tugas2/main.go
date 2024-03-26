package main

import (
	"fmt"
)

func main() {

	var (
		massaMark, tinggiMark float64
		massaJohn, tinggiJohn float64
	)

	fmt.Println("Masukkan massa (kg) Mark dan tinggi (m) Mark:")
	fmt.Scanln(&massaMark, &tinggiMark)
	fmt.Scanln()

	fmt.Println("Masukkan massa (kg) John dan tinggi (m) John:")
	fmt.Scanln(&massaJohn, &tinggiJohn)
	fmt.Scanln()

	bmiMark := hitungBMI(massaMark, tinggiMark)
	bmiJohn := hitungBMI(massaJohn, tinggiJohn)

	markHigherBMI := bmiMark > bmiJohn

	fmt.Printf("BMI Mark: %.2f\n", bmiMark)
	fmt.Printf("BMI John: %.2f\n", bmiJohn)
	if markHigherBMI {
		fmt.Println("Mark memiliki BMI lebih tinggi dari John.")
	} else {
		fmt.Println("John memiliki BMI lebih tinggi dari Mark.")
	}
}

func hitungBMI(massa, tinggi float64) float64 {
	return (massa / (tinggi * tinggi)) * 703
}
