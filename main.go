package main

import (
	"errors"
	"fmt"
)
// Verileri float64 olarak alıyoruz, float64 ve varsa error tipinde geri gönderiyoruz
func topla(birinci, ikinci float64) (float64, error) {
	// Sayılar 0'a eşit veya küçük mü? kontrolü. 
	if birinci <= 0 || ikinci <= 0 {
		return 0, errors.New("Sayılar 0 olamaz")
	}

	// Sorun yoksa, iki değeri topluyoruz, sonuc ve hata kontrolü ile geri gönderiyoruz
	return birinci + ikinci, nil
}

// Verileri float64 olarak alıyoruz, float64 ve varsa error tipinde geri gönderiyoruz
func cikar(birinci, ikinci float64) (float64, error) {
	// Sayılar 0'a eşit veya küçük mü? kontrolü. 
	if birinci <= 0 || ikinci <= 0 {
		return 0, errors.New("Sayılar 0 olamaz")
	}

	// Sorun yoksa, iki değeri çıkartıyoruz, sonuc ve hata kontrolü ile geri gönderiyoruz
	return birinci - ikinci, nil
}

// Verileri float64 olarak alıyoruz, float64 ve varsa error tipinde geri gönderiyoruz
func carp(birinci, ikinci float64) (float64, error) {
	// Sayılar 0'a eşit veya küçük mü? kontrolü. 
	if birinci <= 0 || ikinci <= 0 {
		return 0, errors.New("Sayılar 0 olamaz")
	}
	
	// Sorun yoksa, iki değeri çarpıyoruz, sonuc ve hata kontrolü ile geri gönderiyoruz
	return birinci * ikinci, nil
}

// Verileri float64 olarak alıyoruz, float64 ve varsa error tipinde geri gönderiyoruz
func bol(birinci, ikinci float64) (float64, error) {
	// Sayılar 0'a eşit veya küçük mü? kontrolü. 
	if birinci <= 0 || ikinci <= 0 {
		return 0, errors.New("Sayılar 0 olamaz")
	}

	// Sorun yoksa, iki değeri bölüyoruz, sonuc ve hata kontrolü ile geri gönderiyoruz
	return birinci / ikinci, nil
}

func main() {
	var a, b float64 // Alacağımız değerler için değişken adını ve tipini belirtiyoruz
	var islem string // Kullanıcının hangi işlemi seçtiğiniz string olarak alıyoruz
	var sonuc float64 // Sonucu atayacağımız değişkenimiz
	var err error // Hata için kullanacağımız değişkenimiz.

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

	// İlk Sayı: 4, İkinci Sayı: 5, İşlem: / = "Girilen değerler: 4.00 / 5.00 = 0.80
	fmt.Printf("Girilen değerler: %.2f %s %.2f = %.2f\n", a, islem, b, sonuc)
}
