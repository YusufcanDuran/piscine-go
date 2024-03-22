package main

import (
	"os"
	"strconv"
)

func PrintRune(r rune) {
	os.Stdout.WriteString(string(r))
}

func Atoi(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return num
}

func main() {
	// Programın çağrıldığı komut satırından alınan argüman sayısını kontrol et
	if len(os.Args) != 2 {
		// Eğer argüman sayısı 2 değilse (program adı dahil), hiçbir şey yazdırma
		return
	}

	// Komut satırından alınan argümanı bir tamsayıya çevir
	girilenSayi := Atoi(os.Args[1])

	// Girilen sayının 2'nin üssü olup olmadığını kontrol et
	sonuc := girilenSayi > 0 && (girilenSayi&(girilenSayi-1)) == 0

	// Sonucu ekrana yazdır
	if sonuc {
		PrintRune('t')
		PrintRune('r')
		PrintRune('u')
		PrintRune('e')
	} else {
		PrintRune('f')
		PrintRune('a')
		PrintRune('l')
		PrintRune('s')
		PrintRune('e')
	}
	PrintRune('\n')
}
