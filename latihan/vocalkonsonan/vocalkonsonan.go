package main

import (
	"fmt"
)

func main() {
	isVokal('r')
	isVokal('a')
}

func isVokal(character rune) {
	if character == 'a' || character == 'e' || character == 'i' || character == 'o' || character == 'u' {
		fmt.Printf(" %c adalah vokal\n", character)
	} else {
		fmt.Printf(" %c adalah konsonan\n", character)
	}

}

func isVokalSwitch(character rune) {
	switch character {
	case 'a', 'e', 'i', 'o', 'u':
		fmt.Printf(" %c adalah vokal\n", character)
	default:
		fmt.Printf(" %c adalah konsonan\n", character)
	}
}
