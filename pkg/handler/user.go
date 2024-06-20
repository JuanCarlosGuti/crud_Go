package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/JuanCarlosGuti/Go_users.git/internal/request"
	"github.com/JuanCarlosGuti/Go_users.git/internal/user"
	"github.com/JuanCarlosGuti/Go_users.git/pkg/transport"
	"log"
	"net/http"
)

func NewUserHttpServer(ctx context.Context, router *http.ServeMux, endpoints user.Endpoints) {
	router.HandleFunc("/users/", UserServer(ctx, endpoints))
}

func UserServer(ctx context.Context, endpoints user.Endpoints) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		url := r.URL.Path
		log.Println(r.Method, ": ", url)
		path, pathsize := transport.Clean(url)

		if pathsize < 3 || pathsize > 4 {
			InvalidMethod(w)
			log.Println("Invalid path size: ", path)
			return
		}
		params := make(map[string]string)
		if pathsize == 4 && path[2] != "" {
			params["userId"] = path[2]

		}
		ctx = context.WithValue(ctx, "params", params)
		log.Println("valid path size: ", path)
		tran := transport.New(w, r, ctx)

		switch r.Method {
		case http.MethodGet:
			switch pathsize {
			case 3:
				tran.Server(
					transport.Endpoint(endpoints.GetAll),
					decodeGetAlluser,
					encodeResponse,
					encodeError)
			case 4:
				tran.Server(
					nil,
					decodeGetUser,
					encodeResponse,
					encodeError)
			}
		case http.MethodPost:
			if pathsize == 3 {
				tran.Server(
					transport.Endpoint(endpoints.Create),
					decodeCreateuser,
					encodeResponse,
					encodeError)
			}
		case http.MethodPut:
			if pathsize == 3 {
				tran.Server(
					transport.Endpoint(endpoints.Update),
					decodeUpdateuser,
					encodeResponse,
					encodeError)
			}
		default:
			InvalidMethod(w)
		}
	}
}
func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	status := http.StatusInternalServerError
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	fmt.Fprintf(w, `{"status": %d, "message": "%s"}`, w, err.Error())

}

func decodeGetAlluser(ctx context.Context, r *http.Request) (interface{}, error) {
	return nil, nil

}

func decodeCreateuser(ctx context.Context, r *http.Request) (interface{}, error) {
	var req request.CreateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, fmt.Errorf("invalid request body: '%v'", err.Error())
	}
	return req, nil

}
func decodeGetUser(ctx context.Context, r *http.Request) (interface{}, error) {
	//	params := ctx.Value("params").(map[string]string)
	//fmt.Println(params)
	//fmt.Println(params["userId"])
	return nil, fmt.Errorf("error get User")
}
func decodeUpdateuser(ctx context.Context, r *http.Request) (interface{}, error) {
	var req request.CreatePreviewRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, fmt.Errorf("invalid request body: '%v'", err.Error())
	}
	return req, nil

}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	data, err := json.Marshal(response)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
	return nil

}

func InvalidMethod(w http.ResponseWriter) {
	MsgResponse(w, http.StatusNotFound, "Method does not exist")
}

func MsgResponse(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  status,
		"message": message,
	})
}
