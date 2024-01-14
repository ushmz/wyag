package domain

type GitObject interface {
	Fmt() string
	Serialize() []byte
	Deserialize([]byte)
}
