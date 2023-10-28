package main

import "fmt"

func main() {

	fmt.Println("Merhaba\n\n\nDünya") // \n \t
	fmt.Println("Hello\t\t\tWorld")

	fmt.Println("muhammet alinin \"Dünyası\" ")

	fmt.Println("Sakarya" + "ıslamaköfte" + "Ünlüdür")
	fmt.Println("Malatya", "Kayısısı", "Ünlüdür")

	fmt.Println("Lorem İpsum" +
		"Merhaba Dünya" +
		"Hello World")

	isim := "Muhammet ali"
	soyisim := "özbay"
	sehir := "Sakarya"

	fmt.Println(isim, soyisim, sehir)

	//fmt.Println(len(isim))

	var stringUzunlugu = len(isim)

	fmt.Println(stringUzunlugu)

}
