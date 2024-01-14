package domain

type Tree struct {
	Data []byte
}

func (t *Tree) Fmt() string { return "tree" }
func (t *Tree) Serialize() []byte {
	return t.Data
}
func (t *Tree) Deserialize(data []byte) {
	t.Data = data
}
