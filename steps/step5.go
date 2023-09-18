package steps

import "fmt"

func Step5() {
	if IsBusy {
		fmt.Println("inchid telefonul")
		Step11()
		return
	}

	fmt.Println("astept sa raspunda")
	Step6()
	return
}
