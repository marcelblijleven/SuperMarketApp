package basket

import (
	"errors"
)

type Manager struct {
	repository Repository
}

func (m *Manager) Get(id string) (Basket, error) {
	return m.repository.Get(id)
}

func (m *Manager) Create() (Basket, error) {
	return m.repository.Create()
}

func (m *Manager) Update(basketID string, line ProductLine) (Basket, error) {
	return Basket{}, errors.New("not implemented yet")
}

func (m *Manager) Delete(id string) error {
	return m.repository.Delete(id)
}

func NewManager(repository Repository) Manager {
	return Manager{repository: repository}
}