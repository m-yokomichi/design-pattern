package main

import (
	"./composite"
)

func main() {
	dir := composite.CreateDirectory("ディレクトリ名")

	file1 := composite.CreateFile("ファイル名1")
	file2 := composite.CreateFile("ファイル名2")
	file3 := composite.CreateFile("ファイル名3")

	dir.Add(file1)
	dir.Add(file2)
	dir.Add(file3)

	dir.Remove()
}
