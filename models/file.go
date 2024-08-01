package models

import "gorm.io/gorm"

type File struct {
	gorm.Model
	UID       string `json:"uid"`
	CreatedBy uint   `json:"created_by"` // userID
	Hash      string `json:"hash"`
	Mimetype  string `json:"mimetype"`
	Data      []byte `json:"data"`
}

func (f *File) TableName() string {
	return "files"
}
