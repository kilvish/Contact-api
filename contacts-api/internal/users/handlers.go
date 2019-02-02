package users

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	serr "github.com/anil/contacts-api/error"

	"github.com/anil/contacts-api/middleware"
)

const (
	timeout = 10
)

type handlers struct {
	addUserHandler http.Handler
	getUserHander  http.Handler
}

func getHandlers(s Service) *handlers {
	middlewares := getMiddlewares()
	return &handlers{
		addUserHandler: middlewares.addUserMiddlewares.Then(makeAddUserHandler(s)),
		getUserHander:  middlewares.getUserMiddlewares.Then(makeGetUserHandler(s)),
	}
}

func makeAddUserHandler(s Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx, cancelFunc := context.WithTimeout(r.Context(), time.Duration(timeout*time.Second))
		defer cancelFunc()
		reqID := ctx.Value(middleware.ContextKey(middleware.RequestIDContextKey)).(string)
		req := &addUserRequest{}
		req.load(r)
		resData, err := s.addUser(ctx, req)
		res := new(response)
		res.RequestID = reqID
		if err != nil {
			w.WriteHeader(serr.GetHTTPCode(err))
			res.Error = err.Error()
			res.Status = failureStatus
			fmt.Fprintln(w, res)
			return
		}
		res.responseData = *resData
		res.Status = successStatus
		json.NewEncoder(w).Encode(res)
	})
}

func makeGetUserHandler(s Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx, cancelFunc := context.WithTimeout(r.Context(), time.Duration(timeout*time.Second))
		defer cancelFunc()
		reqID := ctx.Value(middleware.ContextKey(middleware.RequestIDContextKey)).(string)
		req := &getUserRequest{}
		req.load(r)
		resData, err := s.getUser(ctx, req)
		res := new(response)
		res.RequestID = reqID
		if err != nil {
			w.WriteHeader(serr.GetHTTPCode(err))
			res.Error = err.Error()
			res.Status = failureStatus
			fmt.Fprintln(w, res)
			return
		}
		res.responseData = *resData
		res.Status = successStatus
		json.NewEncoder(w).Encode(res)
	})
}
