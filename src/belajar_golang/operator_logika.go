package main

import (
	"fmt"
)

func main() {
	var left = false
	var right = true

	var leftAndRight = left && right
	fmt.Printf("Left dan Right \t(%t) \n", leftAndRight) //Nilainya false karena hasil dari false dan true adalah false

	var leftOrRight = left || right
	fmt.Printf("Left atau Right \t(%t) \n", leftOrRight) //Nilainya true karena hasil dari false atau true adalah true

	var leftReverse = !left
	fmt.Printf("Tidak Left \t\t(%t)\n", leftReverse) //Nilainya true karena hasil dari negasi(atau lawan dari) false adalah true	
}
