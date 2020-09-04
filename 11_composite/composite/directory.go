package composite

type Directory struct {
	list []DirectryEntry
	name string
}

func CreateDirectory(name string) *Directory {
	dir := &Directory{
		name: name,
	}
	return dir
}

func (dir *Directory) Add(entry DirectryEntry) {
	dir.list = append(dir.list, entry)
}

func (dir *Directory) Remove() {
	for _, entry := range dir.list {
		entry.Remove()
	}
}
