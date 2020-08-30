package bridge

import (
	"fmt"
)
type BubbleSortlmple struct {}

func (s *BubbleSortlmple) Sort() {
	// バブルソート実装 (割愛)
	fmt.Println("バブルソートです")
	return 
}

func CreateBubbleSortlmple() Sortlmple {
	var sort Sortlmple
	sort = &BubbleSortlmple{}
	return sort
}