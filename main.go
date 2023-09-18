package main

import (
	"flag"
	"fmt"

	"github.com/andrei-dascalu/go-phone/steps"
)

func main() {
	flag.BoolVar(&steps.HasTone, "ton", false, "telefonul are ton?")
	flag.BoolVar(&steps.IsBusy, "ocupat", false, "telefonul suna ocupat?")
	flag.BoolVar(&steps.Answering, "rasp", false, "raspunde cineva?")
	flag.BoolVar(&steps.PersonAnswered, "pers", false, "persoana cautata a raspuns?")
	flag.BoolVar(&steps.PersonCanCome, "vine", false, "persoana cautata poate veni?")
	flag.IntVar(&steps.WaitMaxAttempts, "incercari", 3, "de cate ori incerc")
	help := flag.Bool("h", false, "afiseaza parametri")
	flag.Parse()

	if help != nil && *help {
		fmt.Println("\nCum vorbeau bunicii la telefon? Parametrii sunt <false>\n\n")
		flag.PrintDefaults()
		return
	}

	steps.Step1()
}
