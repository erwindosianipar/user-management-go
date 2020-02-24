package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"usermanagement/common"
	"usermanagement/models"
	"usermanagement/user"
)

type UserHandler struct {
	userService user.UserService
}

func CreateUserHandler(r *mux.Router, userService user.UserService) {
	userHandler := UserHandler{userService}

	r.HandleFunc("/user/register", userHandler.register).Methods(http.MethodPost)
	r.HandleFunc("/user/login", userHandler.login).Methods(http.MethodPost)
	r.HandleFunc("/user/all", userHandler.getAllUser).Methods(http.MethodGet)
	r.HandleFunc("/user/{id}", userHandler.getUserById).Methods(http.MethodGet)
	r.HandleFunc("/user/{id}", userHandler.updateUser).Methods(http.MethodPut)
	r.HandleFunc("/user/{id}", userHandler.deleteUser).Methods(http.MethodDelete)
}

func (h *UserHandler) register(writer http.ResponseWriter, request *http.Request) {
	reqBody, err := ioutil.ReadAll(request.Body)
	if err != nil {
		common.Response(writer, http.StatusBadRequest, common.Message(false, "Oops, something went wrong"))
		logrus.Error("[UserHandler.register.ioutil.ReadAll] ", err)
		return
	}

	reqUser := models.User{}

	err = json.Unmarshal(reqBody, &reqUser)
	if err != nil {
		common.Response(writer, http.StatusBadRequest, common.Message(false, "Request body is not valid"))
		logrus.Error("[UserHandler.register.json.Unmarshal] ", err)
		return
	}

	common.Response(writer, http.StatusCreated, nil)
	return
}

func (h *UserHandler) login(writer http.ResponseWriter, request *http.Request) {

}

func (h *UserHandler) getAllUser(writer http.ResponseWriter, request *http.Request) {

}

func (h *UserHandler) getUserById(writer http.ResponseWriter, request *http.Request) {

}

func (h *UserHandler) updateUser(writer http.ResponseWriter, request *http.Request) {

}

func (h *UserHandler) deleteUser(writer http.ResponseWriter, request *http.Request) {

}
