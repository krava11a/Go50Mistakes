package main

import "fmt"

func main() {

	//wrong
	isSpace := func(ch byte) bool {
		switch ch {
		case ' ': //mistake
		case '\t':
			return true

		}
		return false
	}

	fmt.Println("Wrong:")
	fmt.Println(isSpace('\t'))
	fmt.Println(isSpace(' '))

	//right
	isSpace = func(ch byte) bool {
		switch ch {
		case ' ','\t':
			return true

		}
		return false
	}

	fmt.Println("Right:")
	fmt.Println(isSpace('\t'))
	fmt.Println(isSpace(' '))

	//fallthrough
	isSpace = func(ch byte) bool {
		switch ch {
		case ' ':
			fallthrough
		case '\t':
			return true

		}
		return false
	}

	fmt.Println("Fallthrough:")
	fmt.Println(isSpace('\t'))
	fmt.Println(isSpace(' '))
}
