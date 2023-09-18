package steps

import "fmt"

func Step4() {
	if HasTone {
		fmt.Println("are ton si formez numarul de telefon")
		Step5()
		return
	}
	fmt.Println("nu are ton - plec la un vecin")
	Step10()
	return
}
