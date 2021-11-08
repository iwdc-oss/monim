package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type UserController interface {
	Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindByUsernameAndPassword(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
