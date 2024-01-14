package domain

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
