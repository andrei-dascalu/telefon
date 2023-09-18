package steps

import "fmt"

func Step7() {
	if PersonAnswered {
		fmt.Println("a raspuns persoana cautata")
		Step9()
		return
	}

	fmt.Println("cer sa vina persoana cautata")
	Step8()
	return
}
