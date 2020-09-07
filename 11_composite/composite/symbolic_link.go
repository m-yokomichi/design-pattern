package composite

import "fmt"

type SymbolicLink struct {
	name string
}

func CreateSymbolicLink(name string) DirectoryEntry{
	var symbolicLink DirectoryEntry
	symbolicLink = &SymbolicLink{name: name}
	return symbolicLink
}

func (s *SymbolicLink) Remove() {
	// 消したことにする
	fmt.Println(s.name, "が削除されました")
}