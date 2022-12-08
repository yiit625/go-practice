package service

import (
	"fmt"
	"github.com/ashishjuyal/banking/domain"
	"io"
	"net/http"
	"os"
	"strings"
)

type FileService interface {
	UploadImage(w http.ResponseWriter, r *http.Request)
}

type DefaultFileService struct {
	repo domain.FileRepository
}

func NewFileService(repository domain.FileRepository) DefaultFileService {
	return DefaultFileService{repository}
}

func (s DefaultFileService) UploadImage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("File Upload Endpoint Hit")

	// Parse our multipart form, 10 << 20 specifies a maximum
	// upload of 10 MB files.
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		fmt.Println("Error parsing multipart form")
		fmt.Println(err)
		return
	}
	// FormFile returns the first file for the given key `myFile`
	// it also returns the FileHeader, so we can get the Filename,
	// the Header and the size of the file
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	// Create a temporary file within our temp-files directory that follows
	// a particular naming pattern
	tempFile, err := os.CreateTemp("temp-files", "upload-*.png")
	if err != nil {
		fmt.Println(err)
	}
	defer tempFile.Close()

	// Write to DB
	if _, err := s.repo.SaveImage(domain.NewFile(
		"",
		"C:\\Users\\User\\Documents\\go-practice\\"+tempFile.Name(),
		strings.Split(tempFile.Name(), "\\")[len(strings.Split(tempFile.Name(), "\\"))-1])); err != nil {
		panic(err)
	}

	// byte array
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	// write this byte array to our temporary file
	tempFile.Write(fileBytes)
	// return that we have successfully uploaded our file!
	fmt.Fprintf(w, "Successfully Uploaded File\n")
}
