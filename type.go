package main

type Commit struct {
	Data []byte
}

func (c *Commit) Fmt() string { return "commit" }
func (c *Commit) Serialize() []byte {
	return c.Data
}
func (c *Commit) Deserialize(data []byte) {
	c.Data = data
}

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

type Tag struct {
	Data []byte
}

func (t *Tag) Fmt() string { return "tag" }
func (t *Tag) Serialize() []byte {
	return t.Data
}
func (t *Tag) Deserialize(data []byte) {
	t.Data = data
}

type Blob struct {
	Data []byte
}

func (b *Blob) Fmt() string { return "blob" }
func (b *Blob) Serialize() []byte {
	return b.Data
}
func (b *Blob) Deserialize(data []byte) {
	b.Data = data
}

type GitObject interface {
	Fmt() string
	Serialize() []byte
	Deserialize([]byte)
}
