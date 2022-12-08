package receiverApp

import (
	"github.com/ashishjuyal/banking/service"
	"net/http"
)

type FileHandler struct {
	service service.FileService
}

func (h FileHandler) NewImage(w http.ResponseWriter, r *http.Request) {
	h.service.UploadImage(w, r)
}
