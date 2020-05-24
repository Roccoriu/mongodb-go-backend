package helpers

import (
	"encoding/json"
	"github.com/roccoriu/restapi/models"
	"log"
	"net/http"
)

func GetError(err error, w http.ResponseWriter) {
	defer log.Println(err.Error())

	response := models.ErrorResponse{
		StatusCode:   http.StatusInternalServerError,
		ErrorMessage: err.Error(),
	}

	message, _ := json.Marshal(response)
	w.WriteHeader(response.StatusCode)

	_, errW := w.Write(message)
	if errW != nil {
		log.Println(errW)
	}
}
