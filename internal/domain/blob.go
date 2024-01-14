package domain

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
