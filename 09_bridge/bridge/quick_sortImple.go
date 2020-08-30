package bridge

import (
	"fmt"
)

type QuickSortlmple struct {}

func (s *QuickSortlmple) Sort() {
	// クイックソート実装 (割愛)
	fmt.Println("クイックソートです")
	return 
}

func CreateQuickSortlmple() Sortlmple {
	var sort Sortlmple
	sort = &QuickSortlmple{}
	return sort
}