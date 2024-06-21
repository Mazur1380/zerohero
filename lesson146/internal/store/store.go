package store

import (
	"github.com/andreipimenov/zerohero/lesson146/internal/model"
	"github.com/google/uuid"
)

type Store struct {
	orders map[uuid.UUID]*model.Order
}

func New() *Store {
	return &Store{
		orders: map[uuid.UUID]*model.Order{},
	}
}

func (s *Store) Push(order *model.Order) {
	s.orders[order.ID] = order
}
func (s *Store) List() map[uuid.UUID]*model.Order {
	return s.orders
}

func (s *Store) Pull(id uuid.UUID) bool {
	_, ok := s.orders[id]
	if !ok {
		return false
	}

	delete(s.orders, id)
	return true
}
