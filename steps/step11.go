package steps

import "fmt"

func Step11() {
	if Waited <= WaitMaxAttempts {
		fmt.Println("astept 15 minute")
		Waited++
		Step2()
		return
	}

	fmt.Println("am asteptat deajuns")
	Step13()
	return
}
