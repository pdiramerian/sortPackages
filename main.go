package main

import "fmt"

const (
	standard = "STANDARD"
	special  = "SPECIAL"
	rejected = "REJECTED"
)

func sort(width int, height int, length int, mass int) string {

	volume := width * height * length

	bulky := volume >= 1_000_000 || width >= 150 || height >= 150 || length >= 150

	isHeavy := mass >= 20

	if bulky && isHeavy {
		return rejected
	}

	if bulky || isHeavy {
		return special
	}

	return standard

}

func main() {
	//STANDARD
	fmt.Println(sort(100, 100, 90, 10))

	// SPECIAL
	fmt.Println(sort(200, 50, 50, 10)) // SPECIAL

	// REJECTED
	fmt.Println(sort(200, 200, 200, 30)) // REJECTED

}
