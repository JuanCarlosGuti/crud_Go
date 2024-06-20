package user

import (
	"context"
	"errors"
	"github.com/JuanCarlosGuti/Go_users.git/internal/domain"
	"github.com/JuanCarlosGuti/Go_users.git/internal/request"
)

type (
	Controller func(ctx context.Context, r interface{}) (interface{}, error)

	Endpoints struct {
		Create Controller
		GetAll Controller
		Update Controller
	}
)

func MakeEndpoints(ctx context.Context, s Service) Endpoints {
	return Endpoints{
		Create: makeCreateEndpoint(s),
		GetAll: makeGetAllEnpoint(s),
		Update: makeUpdateEndpoints(s),
	}
}

func makeGetAllEnpoint(s Service) Controller {
	return func(ctx context.Context, r interface{}) (interface{}, error) {
		users, err := s.GetAll(ctx)
		if err != nil {
			return nil, errors.New("Error getting all users")
		}
		return users, nil

	}

}

func makeCreateEndpoint(s Service) Controller {
	return func(ctx context.Context, r interface{}) (interface{}, error) {
		req := r.(request.CreateRequest)

		if req.FirstName == "" || req.LastName == "" || req.Email == "" {
			//	MsgResponse(w, http.StatusBadRequest, "All fields must be filled")
			return nil, errors.New("first name, last name, email required")
		}
		user, err := s.Create(ctx, req.FirstName, req.LastName, req.Email)
		if err != nil {
			return nil, err
		}
		return user, nil
	}
}

func makeUpdateEndpoints(s Service) Controller {
	return func(ctx context.Context, r interface{}) (interface{}, error) {
		req := r.(request.CreatePreviewRequest)
		if req.FirstName == "" || req.LastName == "" || req.Email == "" {
			return nil, errors.New("first name, last name, email required")
		}

		user := &domain.User{
			ID:        req.ID,
			FirstName: req.FirstName,
			LastName:  req.LastName,
			Email:     req.Email,
		}

		updatedUser, err := s.Update(ctx, user)
		if err != nil {
			return nil, err
		}
		return updatedUser, nil
	}

}
