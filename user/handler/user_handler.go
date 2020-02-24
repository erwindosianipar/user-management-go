package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"usermanagement/commons"
	"usermanagement/middleware"
	"usermanagement/models"
	"usermanagement/user"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	userService user.UserService
}

func CreateUserHandler(r *mux.Router, userService user.UserService) {
	userHandler := UserHandler{userService}


	r.HandleFunc("/register", userHandler.register).Methods(http.MethodPost)
	r.HandleFunc("/login", userHandler.login).Methods(http.MethodPost)

	auth := r.NewRoute().Subrouter()
	auth.HandleFunc("/user/all", middleware.TokenVerifyMiddleware(userHandler.getAllUser)).Methods(http.MethodGet)
	auth.HandleFunc("/user/{id}", middleware.TokenVerifyMiddleware(userHandler.getUserById)).Methods(http.MethodGet)
	auth.HandleFunc("/user/{id}", middleware.TokenVerifyMiddleware(userHandler.updateUser)).Methods(http.MethodPut)
	auth.HandleFunc("/user/{id}", middleware.TokenVerifyMiddleware(userHandler.deleteUser)).Methods(http.MethodDelete)
}

func (h *UserHandler) register(writer http.ResponseWriter, request *http.Request) {
	reqBody, err := ioutil.ReadAll(request.Body)
	if err != nil {
		commons.Response(writer, http.StatusBadRequest, commons.Message(false, "Oops, something went wrong."))
		logrus.Error(err)
		return
	}

	reqUser := models.User{}
	err = json.Unmarshal(reqBody, &reqUser)
	if err != nil {
		commons.Response(writer, http.StatusBadRequest, commons.Message(false, "Request body is not valid."))
		logrus.Error(err)
		return
	}

	reqRegisterValid := commons.ValidateNameEmailPass(writer, reqUser)
	if !(reqRegisterValid) {
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(reqUser.Password), 10)
	if err != nil {
		commons.Response(writer, http.StatusInternalServerError, commons.Message(false, "Oops, something went wrong. Please try again."))
		logrus.Error(err)
		return
	}

	reqUser.Password = string(hashedPassword)
	newUser, err := h.userService.Register(&reqUser)
	if err != nil {
		commons.Response(writer, http.StatusInternalServerError, commons.Message(false, err.Error()))
		logrus.Error(err)
		return
	}

	commons.Response(writer, http.StatusCreated, newUser)
	return
}

func (h *UserHandler) login(writer http.ResponseWriter, request *http.Request) {
	var jwt models.JWT
	reqBody, err := ioutil.ReadAll(request.Body)
	if err != nil {
		commons.Response(writer, http.StatusBadRequest, commons.Message(false, "Oops, something went wrong."))
		logrus.Error(err)
		return
	}

	reqUser := models.User{}
	err = json.Unmarshal(reqBody, &reqUser)
	if err != nil {
		commons.Response(writer, http.StatusBadRequest, commons.Message(false, "Request body is not valid."))
		logrus.Error(err)
		return
	}

	reqLoginValid := commons.ValidateEmailPass(writer, reqUser)
	if !(reqLoginValid) {
		return
	}

	dataUser, err := h.userService.GetUserByEmail(reqUser.Email)
	if err != nil {
		commons.Response(writer, http.StatusBadRequest, commons.Message(false, err.Error()))
		logrus.Error(err)
		return
	}

	hashedPassword := dataUser.Password
	if !(commons.ComparePassword(hashedPassword, []byte(reqUser.Password))) {
		commons.Response(writer, http.StatusBadRequest, commons.Message(false, "Email or password did not match."))
		logrus.Error(err)
		return
	}

	token, err := commons.GenerateToken(dataUser)
	if err != nil {
		logrus.Fatal("[UserHandler.login.GenerateToken]", err)
	}

	jwt.Token = token
	jwt.Id = dataUser.ID

	mapResponse := commons.Message(true, "success: user already logged in")
	mapResponse["response"] = jwt

	commons.Response(writer, http.StatusOK, mapResponse)
	return
}

func (h *UserHandler) getAllUser(writer http.ResponseWriter, request *http.Request) {
	response, err := h.userService.GetAllUser()
	if err != nil {
		commons.Response(writer, http.StatusBadRequest, commons.Message(false, err.Error()))
		return
	}

	commons.Response(writer, http.StatusOK, response)
	return
}

func (h *UserHandler) getUserById(writer http.ResponseWriter, request *http.Request) {
	param := mux.Vars(request)
	id, err := strconv.Atoi(param["id"])
	if err != nil {
		commons.Response(writer, http.StatusBadRequest, commons.Message(false, "Please insert a valid number."))
		return
	}

	response, err := h.userService.GetUserByID(uint(id))
	if err != nil {
		commons.Response(writer, http.StatusBadRequest, commons.Message(false, err.Error()))
		return
	}

	commons.Response(writer, http.StatusOK, response)
	return
}

func (h *UserHandler) updateUser(writer http.ResponseWriter, request *http.Request) {
	dataUser := new(models.User)
	param := mux.Vars(request)
	id, err := strconv.Atoi(param["id"])
	if err != nil {
		commons.Response(writer, http.StatusBadRequest, commons.Message(false, "Please insert a valid number."))
		return
	}

	if err := json.NewDecoder(request.Body).Decode(&dataUser); err != nil {
		commons.Response(writer, http.StatusBadRequest, commons.Message(false, "Invalid request: "+err.Error()))
		return
	}

	response, err := h.userService.UpdateUser(uint(id), dataUser)
	if err != nil {
		commons.Response(writer, http.StatusBadRequest, commons.Message(false, err.Error()))
		return
	}

	commons.Response(writer, http.StatusOK, response)
	return
}

func (h *UserHandler) deleteUser(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		commons.Response(writer, http.StatusBadRequest, commons.Message(false, "Please insert a valid number."))
		return
	}

	response, err := h.userService.DeleteUser(uint(id))
	if err != nil {
		commons.Response(writer, http.StatusBadRequest, commons.Message(false, err.Error()))
		return
	}

	commons.Response(writer, http.StatusOK, response)
	return
}
