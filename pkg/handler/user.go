package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/JuanCarlosGuti/Go_users.git/internal/request"
	"github.com/JuanCarlosGuti/Go_users.git/internal/user"
	"github.com/JuanCarlosGuti/Go_users.git/pkg/transport"
	response2 "github.com/JuanCarlosGuti/Response_GO/response"
	"log"
	"net/http"
	"strconv"
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

		tran := transport.New(w, r, context.WithValue(ctx, "params", params))
		var end user.Controller
		var deco func(ctx context.Context, r *http.Request) (interface{}, error)

		switch r.Method {
		case http.MethodGet:
			switch pathsize {
			case 3:

				end = endpoints.GetAll
				deco = decodeGetAlluser

			case 4:
				end = endpoints.Get
				deco = decodeGetUser

			}
		case http.MethodPost:
			switch pathsize {
			case 3:
				end = endpoints.Create
				deco = decodeCreateuser

			}
		case http.MethodPut:
			if pathsize == 3 {
				end = endpoints.Update
				deco = decodeUpdateuser

			}
		case http.MethodPatch:
			switch pathsize {
			case 4:
				end = endpoints.UpdateP
				deco = decodeUpdateUser2
			}

		}
		if end != nil && deco != nil {
			tran.Server(
				transport.Endpoint(end),
				deco,
				encodeResponse,
				encodeError)
		} else {
			InvalidMethod(w)
		}

	}
}
func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	resp := err.(response2.Response)

	w.WriteHeader(resp.StatusCode())
	json.NewEncoder(w).Encode(resp)

}

func decodeGetUser(ctx context.Context, r *http.Request) (interface{}, error) {
	params := ctx.Value("params").(map[string]string)
	id, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
		return nil, err
	}

	return user.GetReq{
		ID: id,
	}, nil

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
func decodeUpdateUser2(ctx context.Context, r *http.Request) (interface{}, error) {
	var req request.UpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, fmt.Errorf("invalid request body: '%v'", err.Error())
	}

	params := ctx.Value("params").(map[string]string)
	id, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
		return nil, err
	}
	log.Println("id del params: ", id)
	req.ID = id
	return req, nil
}

func decodeUpdateuser(ctx context.Context, r *http.Request) (interface{}, error) {
	var req request.CreatePreviewRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, fmt.Errorf("invalid request body: '%v'", err.Error())
	}
	return req, nil

}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	r := response.(response2.Response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.StatusCode())

	return json.NewEncoder(w).Encode(r)

}

func InvalidMethod(w http.ResponseWriter) {
	status := http.StatusNotFound
	w.WriteHeader(status)
	fmt.Fprintf(w, `{"status": %d, "message": "methos doesn't exist"}`, status)
}
