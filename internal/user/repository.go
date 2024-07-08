package user

import (
	"context"
	"fmt"
	"github.com/JuanCarlosGuti/Go_users.git/internal/domain"
	"log"
	"slices"
)

type DB struct {
	Users     []domain.User
	MaxUserID uint64
}

type (
	Repository interface {
		Create(ctx context.Context, user *domain.User) error
		GetAll(ctx context.Context) ([]domain.User, error)
		Get(ctx context.Context, id uint64) (*domain.User, error)
		Update(ctx context.Context, user *domain.User) error
		Update2(ctx context.Context, id uint64, firstName *string, lastName *string, email *string) error
	}

	repository struct {
		db  DB
		log *log.Logger
	}
)

func NewRepository(db DB, l *log.Logger) Repository {
	return &repository{
		db:  db,
		log: l,
	}
}
func (r *repository) Create(ctx context.Context, user *domain.User) error {
	r.db.MaxUserID++
	user.ID = r.db.MaxUserID
	r.db.Users = append(r.db.Users, *user)
	r.log.Println("respository.Create", user)
	return nil

}
func (r *repository) GetAll(ctx context.Context) ([]domain.User, error) {
	r.log.Println("repository.GetAll")
	return r.db.Users, nil
}
func (r *repository) Get(ctx context.Context, id uint64) (*domain.User, error) {
	index := slices.IndexFunc(r.db.Users, func(v domain.User) bool {
		return v.ID == id
	})
	if index == -1 {
		return nil, fmt.Errorf("user not found")
	}
	return &r.db.Users[index], nil

}

func (r *repository) Update(ctx context.Context, user *domain.User) error {
	r.log.Println("repository.Update Attempt", user)
	found := false
	for i, u := range r.db.Users {
		if u.ID == user.ID {
			r.db.Users[i] = *user
			r.log.Println("repository.Update Successful", user)
			found = true
			break
		}
	}
	if !found {
		err := fmt.Errorf("no user found with ID %d", user.ID)
		r.log.Println(err)
		return err
	}
	return nil
}

func (r *repository) Update2(ctx context.Context, id uint64, firstName *string, lastName *string, email *string) error {
	user, err := r.Get(ctx, id)
	//log.Println("el Id: ", id)
	//
	//log.Println("user", user, err)
	if err != nil {
		log.Println("Error updating user", err)
		return err
	}
	if firstName != nil && user.FirstName != *firstName {
		user.FirstName = *firstName
	}
	if lastName != nil && user.LastName != *lastName {
		user.LastName = *lastName
	}
	if email != nil {
		user.Email = *email
	}

	return nil

}

//func (r *repository) Update(ctx context.Context, user *domain.User) error {
//	r.log.Println("repository.Update", user)
//	users := r.db.Users
//	for i, u := range users {
//		if u.ID == user.ID {
//			users[i] = *user
//			r.db.Users = users
//			r.log.Println("repository.Update", users[i])
//		}
//	}
//	return nil
//}
