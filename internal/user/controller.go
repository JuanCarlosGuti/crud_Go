package user

import (
	"context"
	"errors"
	"github.com/JuanCarlosGuti/Go_users.git/internal/domain"
	"github.com/JuanCarlosGuti/Go_users.git/internal/request"
	"log"
)

type (
	Controller func(ctx context.Context, r interface{}) (interface{}, error)

	Endpoints struct {
		Create  Controller
		GetAll  Controller
		Get     Controller
		Update  Controller
		UpdateP Controller
	}
	GetReq struct {
		ID uint64
	}
)

func MakeEndpoints(ctx context.Context, s Service) Endpoints {
	return Endpoints{
		Create:  makeCreateEndpoint(s),
		GetAll:  makeGetAllEnpoint(s),
		Get:     makeGetEnpoint(s),
		Update:  makeUpdateEndpoints(s),
		UpdateP: makeUpdateEndpoints2(s),
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
func makeGetEnpoint(s Service) Controller {
	return func(ctx context.Context, r interface{}) (interface{}, error) {
		req := r.(GetReq)

		user, err := s.Get(ctx, req.ID)
		if err != nil {
			return nil, errors.New("Error getting all users")
		}
		return user, nil

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
func makeUpdateEndpoints2(s Service) Controller {
	return func(ctx context.Context, r interface{}) (interface{}, error) {
		req := r.(request.UpdateRequest)
		log.Println("id controller", req.ID)
		if err := s.Update2(ctx, req.ID, req.FirstName, req.LastName, req.Email); err != nil {

			return nil, err
		}
		return nil, nil

	}
}
