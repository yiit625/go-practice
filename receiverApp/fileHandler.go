package receiverApp

import (
	"encoding/json"
	"github.com/ashishjuyal/banking/service"
	"net/http"
)

type FileHandler struct {
	service service.FileService
}

func (h FileHandler) NewImage(w http.ResponseWriter, r *http.Request) {
	fileId, appError := h.service.UploadDocument(service.WriteImage(w, r))
	if appError != nil {
		writeResponse(w, appError.Code, appError.AsMessage())
	} else {
		writeResponse(w, http.StatusCreated, fileId)
	}
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}
