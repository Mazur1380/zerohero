package store

import (
	"errors"
	"fmt"

	"github.com/andreipimenov/zerohero/lesson141/model"
)

type Store struct {
	users    []*model.User
	products []*model.Product
	orders   []*model.Order
}

func New() *Store {
	s := &Store{
		users:    []*model.User{},
		products: []*model.Product{},
		orders:   []*model.Order{},
	}
	return s
}

func (s *Store) CreateUser(id int, name string, balance int) {
	u := &model.User{
		ID:      id,
		Name:    name,
		Balance: balance,
	}
	s.users = append(s.users, u)
}

func (s *Store) GetUserById(id int) (*model.User, error) {
	for _, u := range s.users {
		if u.ID == id {
			return u, nil
		}
	}
	return nil, errors.New("user not found")
}

func (s *Store) CreateProduct(id int, name string, price int) {
	p := &model.Product{
		ID:    id,
		Name:  name,
		Price: price,
	}
	s.products = append(s.products, p)
}

func (s *Store) GetProductById(id int) (*model.Product, error) {
	for _, p := range s.products {
		if p.ID == id {
			return p, nil
		}
	}
	return nil, errors.New("product not found")
}

func (s *Store) CreateOrder(id int, userId int, productId int) error {
	u, err := s.GetUserById(userId)
	if err != nil {
		return err
	}
	p, err := s.GetProductById(productId)
	if err != nil {
		return err
	}
	if u.Balance < p.Price {
		return errors.New("not enough money")
	}
	err = s.SetBalance(userId, u.Balance-p.Price)
	if err != nil {
		return err
	}
	o := &model.Order{
		ID:        id,
		UserId:    userId,
		ProductId: productId,
	}
	s.orders = append(s.orders, o)
	return nil
}

func (s *Store) GetOrderById(id int) (*model.Order, error) {
	for _, o := range s.orders {
		if o.ID == id {
			return o, nil
		}
	}
	return nil, errors.New("order not found")
}

func (s *Store) SetBalance(userId int, newBalance int) error {
	u, err := s.GetUserById(userId)
	if err != nil {
		return err
	}
	u.Balance = newBalance
	return nil
}

func (s *Store) GetOrderHistory(userId int) ([]string, error) {
	res := []string{}
	u, err := s.GetUserById(userId)
	if err != nil {
		return nil, err
	}
	for _, o := range s.orders {
		if o.UserId != userId {
			continue
		}
		p, err := s.GetProductById(o.ProductId)
		if err != nil {
			return nil, err
		}
		s := fmt.Sprintf("user %s bought %s for %d", u.Name, p.Name, p.Price)
		res = append(res, s)
	}
	return res, nil
}
