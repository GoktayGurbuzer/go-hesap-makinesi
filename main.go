package main

import (
	"errors"
	"fmt"
)

func topla(birinci, ikinci float64) (float64, error) {
	if birinci <= 0 || ikinci <= 0 {
		return 0, errors.New("Sayılar 0 olamaz")
	}

	return birinci + ikinci, nil
}

func cikar(birinci, ikinci float64) (float64, error) {
	if birinci <= 0 || ikinci <= 0 {
		return 0, errors.New("Sayılar 0 olamaz")
	}

	return birinci - ikinci, nil
}

func carp(birinci, ikinci float64) (float64, error) {
	if birinci <= 0 || ikinci <= 0 {
		return 0, errors.New("Sayılar 0 olamaz")
	}
	return birinci * ikinci, nil
}

func bol(birinci, ikinci float64) (float64, error) {
	if birinci <= 0 || ikinci <= 0 {
		return 0, errors.New("Sayılar 0 olamaz")
	}

	return birinci / ikinci, nil
}

func main() {
	var a, b float64
	var islem string
	var sonuc float64
	var err error

	fmt.Println("İlk Sayıyı Girin ")
	fmt.Scanln(&a)

	fmt.Println("İkinci Sayıyı Girin ")
	fmt.Scanln(&b)

	fmt.Println("İşlem seçin (+, -, *, /):")
	fmt.Scanln(&islem)

	switch islem {
	case "+":
		sonuc, err = topla(a, b)
	case "-":
		sonuc, err = cikar(a, b)
	case "*":
		sonuc, err = carp(a, b)
	case "/":
		sonuc, err = bol(a, b)
	default:
		fmt.Println("Geçersiz işlem!")
		return
	}

	if err != nil {
		fmt.Println("Hata:", err)
		return
	}

	fmt.Printf("Girilen değerler: %.2f %s %.2f = %.2f\n", a, islem, b, sonuc)
}
