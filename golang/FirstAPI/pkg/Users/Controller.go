package Users

import (
	"encoding/json"
	"net/http"
)

type Controller struct {
	UserRepository *Repository
}

func GetNewUsersController(userRepository *Repository) *Controller {
	return &Controller{
		UserRepository: userRepository,
	}
}

func (controller Controller) Handle(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(controller.UserRepository.GetAll())
}