package models

type Model interface {
	SetId(id string)
	GetId() string
}

type modelImpl struct {
	id   string
	name string
}

func (m *modelImpl) SetId(id string) {
	m.id = id
}
