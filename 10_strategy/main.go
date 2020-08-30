package main

import (
	"fmt"

	"./strategy"
)

func main() {
	var ageCompare strategy.Comparator
	ageCompare = &strategy.AgeComparator{}
	var weightCompare strategy.Comparator
	weightCompare = &strategy.WeightComparator{}
	human1 := strategy.Human{
		Age:    20,
		Weight: 50,
		Height: 160,
	}
	human2 := strategy.Human{
		Age:    18,
		Weight: 60,
		Height: 180,
	}

	fmt.Println(ageCompare.Compare(human1, human2))
	fmt.Println(weightCompare.Compare(human1, human2))
}
