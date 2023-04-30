package models

type Tag struct {
	id   int    `json:"id"`
	name string `json:"name"`
}

func NewTag(name string) (*Tag, error) {
	return &Tag{name: name}, nil
}

func (t *Tag) GetId() int {
	return t.id
}

func (t *Tag) GetName() string {
	return t.name
}

func (t *Tag) SetName(name string) {
	t.name = name
}
