package users

import "github.com/anil/contacts-api/middleware"

type middlewares struct {
	addUserMiddlewares *middleware.Chain
	getUserMiddlewares *middleware.Chain
}

func getMiddlewares() *middlewares {
	return &middlewares{
		addUserMiddlewares: getAddUserMiddleware(),
		getUserMiddlewares: getGetUserMiddleware(),
	}
}

func getAddUserMiddleware() *middleware.Chain {
	return middleware.New(middleware.GetRequestIDMiddleware(), middleware.GetContentTypeMiddleware())
}

func getGetUserMiddleware() *middleware.Chain {
	return middleware.New(middleware.GetRequestIDMiddleware(), middleware.GetContentTypeMiddleware())
}

