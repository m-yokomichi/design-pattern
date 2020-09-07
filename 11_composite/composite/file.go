package composite

import (
	"fmt"
)

type File struct {
	name string
}

func (f *File) Remove() {
	// 消したことにする
	fmt.Println(f.name, "が削除されました")
}

func CreateFile(name string) DirectoryEntry {
	var file DirectoryEntry
	file = &File{name: name}
	return file
}
