package steps

import "fmt"

func Step12() {
	//ca sa nu avem loop infinit daca nu raspunde
	if Waited <= WaitMaxAttempts {
		fmt.Println("astept o ora")
		Waited++
		Step2()
		return
	}

	fmt.Println("am asteptat deajuns")
	Step13()
	return
}
