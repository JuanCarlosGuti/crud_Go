package user

import (
	"context"
	"github.com/JuanCarlosGuti/Go_users.git/internal/domain"
	"log"
)

type (
	Service interface {
		Create(ctx context.Context, firstname, lastname, email string) (*domain.User, error)
		GetAll(ctx context.Context) ([]domain.User, error)
		Update(ctx context.Context, user *domain.User) (*domain.User, error)
	}
	service struct {
		log  *log.Logger
		repo Repository
	}
)

func NewService(l *log.Logger, repo Repository) Service {
	return &service{
		log:  l,
		repo: repo,
	}
}

func (s service) Create(ctx context.Context, firstname, lastname, email string) (*domain.User, error) {
	user := &domain.User{
		FirstName: firstname,
		LastName:  lastname,
		Email:     email,
	}
	err := s.repo.Create(ctx, user)
	if err != nil {
		return nil, err
	}
	s.log.Println("Created user", user)
	return user, nil

}

func (s service) GetAll(ctx context.Context) ([]domain.User, error) {
	users, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return users, nil
}
func (s *service) Update(ctx context.Context, user *domain.User) (*domain.User, error) {
	s.log.Println("Updating user", user)
	err := s.repo.Update(ctx, user)
	if err != nil {
		s.log.Println("Error updating user", err)
		return nil, err // Devuelve nil y el error si hay uno
	}
	s.log.Println("Updated user successfully", user)
	return user, nil // Devuelve el usuario actualizado y nil como error
}
