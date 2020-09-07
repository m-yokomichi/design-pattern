package composite

import "fmt"

type Directory struct {
	DirectoryEntry
	list []DirectoryEntry
	name string
}

func CreateDirectory(name string) *Directory {
	dir := &Directory{
		name:           name,
		DirectoryEntry: &Directory{},
	}
	return dir
}

func (dir *Directory) Add(entry DirectoryEntry) {
	dir.list = append(dir.list, entry)
}

func (dir *Directory) Remove() {
	for _, entry := range dir.list {
		entry.Remove()
	}
	fmt.Println(dir.name, "が削除されました")
}
