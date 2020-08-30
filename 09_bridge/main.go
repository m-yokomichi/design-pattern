package main

import (
	"./bridge"
)

func main() {

	cSort := bridge.CreateSort(bridge.CreateQuickSortlmple())
	bSort := bridge.CreateSort(bridge.CreateBubbleSortlmple())
	cSort.Sortlmple.Sort()
	bSort.Sortlmple.Sort()
}