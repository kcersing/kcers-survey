package utils

import "fmt"

type Tree struct {
	Id    int64
	Tree  string
	Level int64
}

func (t *Tree) prefixTreeKey() string {
	return fmt.Sprintf("tr_%d", t.Id) + " "
}
func (t *Tree) defaultTreeKey() string {
	return "tr_0 "
}

func (t *Tree) GetTree() string {

	return t.Tree + t.prefixTreeKey()
}
func (t *Tree) GetLevel() int64 {

	return t.Level + 1

}
