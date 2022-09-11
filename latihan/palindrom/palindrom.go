package main

import "fmt"

func main() {
	// fmt.Println("PALINDROM")

	// fmt.Println(isPalindromWithTemp("a"))
	// fmt.Println(isPalindromWithTemp("ab"))
	// fmt.Println(isPalindromWithTemp("aba"))
	// fmt.Println(isPalindromWithTemp("katak"))

	fmt.Println(isPalindrom("a"))
	fmt.Println(isPalindrom("ab"))
	fmt.Println(isPalindrom("aba"))
	fmt.Println(isPalindrom("katak"))
}

func isPalindromWithTemp(kata string) bool {
	temp := ""
	for i := len(kata) - 1; i >= 0; i-- {
		// fmt.Println(kata[i])
		temp = temp + string(kata[i])
		// fmt.Print(temp)
	}
	return temp == kata
}

func isPalindrom(kata string) bool {
	// fmt.Print("KATA PALINDROM: ", kata)
	for i := 0; i < len(kata)/2; i++ {
		indexAwal := i
		indexAkhir := len(kata) - i - 1
		// fmt.Println("Awal", indexAwal)
		// fmt.Println(kata[indexAwal], " ", kata[indexAkhir])
		fmt.Println(string(kata[indexAwal]), " ", string(kata[indexAkhir]))

		if kata[indexAwal] != kata[indexAkhir] {
			return false
		}

	}
	return true
}
