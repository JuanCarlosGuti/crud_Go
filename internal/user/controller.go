package user

import (
	"context"
	"errors"
	"github.com/JuanCarlosGuti/Go_users.git/internal/domain"
	"github.com/JuanCarlosGuti/Go_users.git/internal/request"
	"github.com/JuanCarlosGuti/Response_GO/response"
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
			return nil, response.InternalServerError("Error getting all users")
		}
		return response.Ok("successs", users), nil

	}

}
func makeGetEnpoint(s Service) Controller {
	return func(ctx context.Context, r interface{}) (interface{}, error) {
		req := r.(GetReq)

		user, err := s.Get(ctx, req.ID)
		if err != nil {
			if errors.As(err, &ErrorNotFound{}) {
				return nil, response.NotFound(err.Error())
			}
			return nil, ErrorNotFound{req.ID}
		}
		return response.Ok("successs", user), nil

	}

}

func makeCreateEndpoint(s Service) Controller {
	return func(ctx context.Context, r interface{}) (interface{}, error) {
		req := r.(request.CreateRequest)

		if req.FirstName == "" {

			return nil, response.BadRequest(ErrFirstNameRequired.Error())
		}
		if req.LastName == "" {
			return nil, response.BadRequest(ErrLastNameRequired.Error())
		}
		user, err := s.Create(ctx, req.FirstName, req.LastName, req.Email)
		if err != nil {
			return nil, response.InternalServerError(err.Error())
		}
		return response.Crreated("success", user), nil
	}
}

func makeUpdateEndpoints(s Service) Controller {
	return func(ctx context.Context, r interface{}) (interface{}, error) {
		req := r.(request.CreatePreviewRequest)
		if req.FirstName == "" {
			return nil, response.BadRequest(ErrFirstNameRequired.Error())
		}
		if req.LastName == "" {
			return nil, response.BadRequest(ErrLastNameRequired.Error())
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
		if req.FirstName != nil && *req.FirstName == "" {
			return nil, response.BadRequest(ErrFirstNameRequired.Error())
		}
		if req.LastName != nil && *req.LastName == "" {
			return nil, response.BadRequest(ErrLastNameRequired.Error())
		}
		log.Println("id controller", req.ID)
		if err := s.Update2(ctx, req.ID, req.FirstName, req.LastName, req.Email); err != nil {
			if errors.As(err, &ErrorNotFound{}) {
				return nil, response.NotFound(err.Error())
			}
			return nil, response.InternalServerError(err.Error())
		}
		return response.Ok("success", nil), nil

	}
}
