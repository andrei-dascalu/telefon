package steps

import "fmt"

func Step6() {
	if !Answering {
		fmt.Println("nu raspunde nimeni")
		Step12()
		return
	}

	fmt.Println("discut cu cine a raspuns")
	Step7()
	return
}
