package keeper

import (
	"errors"
	"io"
	"mime/multipart"
	"os"
)

func RegisterFile(req *multipart.FileHeader) (err error) {
	fileWriter, err := os.OpenFile("storage/"+req.Filename, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	defer fileWriter.Close()

	fileReader, err := req.Open()
	if err != nil {
		return err
	}
	defer fileReader.Close()

	_, err = io.Copy(fileWriter, fileReader)
	if err != nil {
		return err
	}

	return
}

func RemoveFile(req *string) (err error) {
	if err = os.Remove("storage/"+*req); err != nil {
		return errors.New("file not find")
	}
	return
}
