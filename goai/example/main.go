package main

import (
	"fmt"

	"github.com/needkopi/riset-go/goai"
)

func main() {

	ai := goai.New(goai.Config{
		IsLower:       true,
		IsBinary:      false,
		Smoot:         true,
		NormalizeType: goai.EuclideanSumSquare,
	})
	ai.Learn(PopulateData())

	dataTest := []string{
		"pak tiket ke jakarta ada?",
		"isiin pulsa aku dong",
		"saldo wallet ku tingal dikit nih",
	}

	predicted, err := ai.Predict(dataTest)
	if err != nil {
		panic(err)
	}

	for idx, p := range predicted {
		fmt.Println(dataTest[idx], " >> ", p)
	}
}
