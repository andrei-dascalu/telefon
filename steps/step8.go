package steps

import "fmt"

func Step8() {
	if !PersonCanCome {
		fmt.Println("persoana cautata nu poate veni")
		Step13()
		return
	}

	Step9()
	return
}
