package user

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/JuanCarlosGuti/Go_users.git/internal/domain"
	"github.com/JuanCarlosGuti/Go_users.git/internal/request"
)

type Controller func(w http.ResponseWriter, r *http.Request)
type Endpoints struct {
	Create Controller
	GetAll Controller
	Update Controller
}

func MakeEndpoints(ctx context.Context, s Service) Controller {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			GetAllUser(ctx, s, w)
		case http.MethodPost:
			var req request.CreateRequest
			if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
				MsgResponse(w, http.StatusInternalServerError, err.Error())
				return
			}
			PostUser(ctx, s, w, req)
		case http.MethodPut:
			var req request.CreatePreviewRequest
			if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
				MsgResponse(w, http.StatusInternalServerError, err.Error())
				return
			}
			UpdateUser(ctx, s, w, req)
		default:
			InvalidMethod(w)
		}
	}
}

func InvalidMethod(w http.ResponseWriter) {
	MsgResponse(w, http.StatusNotFound, "Method does not exist")
}

func GetAllUser(ctx context.Context, s Service, w http.ResponseWriter) {
	users, err := s.GetAll(ctx)
	if err != nil {
		MsgResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	DataResponse(w, http.StatusOK, users)
}

func MsgResponse(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  status,
		"message": message,
	})
}

func PostUser(ctx context.Context, s Service, w http.ResponseWriter, req request.CreateRequest) {
	if req.FirstName == "" || req.LastName == "" || req.Email == "" {
		MsgResponse(w, http.StatusBadRequest, "All fields must be filled")
		return
	}
	user, err := s.Create(ctx, req.FirstName, req.LastName, req.Email)
	if err != nil {
		MsgResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	DataResponse(w, http.StatusCreated, user)
}

func UpdateUser(ctx context.Context, s Service, w http.ResponseWriter, data request.CreatePreviewRequest) {
	if data.FirstName == "" || data.LastName == "" || data.Email == "" {
		MsgResponse(w, http.StatusBadRequest, "All fields must be filled")
		return
	}

	user := &domain.User{
		ID:        data.ID,
		FirstName: data.FirstName,
		LastName:  data.LastName,
		Email:     data.Email,
	}

	updatedUser, err := s.Update(ctx, user)
	if err != nil {
		MsgResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	DataResponse(w, http.StatusOK, updatedUser)
}

func DataResponse(w http.ResponseWriter, status int, data interface{}) {
	value, err := json.Marshal(data)
	if err != nil {
		MsgResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	w.WriteHeader(status)
	fmt.Fprintf(w, `{"status": %d, "data": %s}`, status, value)
}
