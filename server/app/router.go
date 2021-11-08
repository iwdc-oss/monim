package app

import (
	"github.com/iwdc-oss/monim/server/controller"
	"github.com/julienschmidt/httprouter"
)

func NewRouter(userController controller.UserController) *httprouter.Router {
	router := httprouter.New()

	router.POST("/api/users", userController.Create)

	return router
}
