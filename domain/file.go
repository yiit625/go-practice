package domain

import "github.com/ashishjuyal/banking-lib/errs"

type File struct {
	Id   string
	Path string
	Name string
}

type FileRepository interface {
	SaveImage(file File) (*File, *errs.AppError)
}
