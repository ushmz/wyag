package domain

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
