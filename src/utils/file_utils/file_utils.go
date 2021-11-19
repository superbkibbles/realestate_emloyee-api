package file_utils

import (
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/superbkibbles/bookstore_utils-go/rest_errors"
	"github.com/superbkibbles/realestate_employee-api/src/utils/crypto_utils"
)

// Save
// If user not saved
// Delete Pic

func DeleteFile(fileName string) {
	if err := os.Remove(filepath.Join("clients/visuals/", filepath.Base(fileName))); err != nil {
		log.Println(err)
	}
}

func SaveFile(header *multipart.FileHeader, file multipart.File) (string, rest_errors.RestErr) {

	// Check if file is Pic Or Video
	splitter := strings.Split(header.Filename, ".")
	ext := splitter[len(splitter)-1]
	fileName := crypto_utils.GetMd5(header.Filename+strconv.FormatInt(time.Now().Unix(), 36)) + "." + ext

	_, err := os.Stat(filepath.Join("clients/visuals/", filepath.Base(fileName)))
	if os.IsNotExist(err) {
		out, err := os.Create(filepath.Join("clients/visuals/", filepath.Base(fileName)))
		if err != nil {
			return "", rest_errors.NewInternalServerErr("Error while creating file", err)
		}

		defer out.Close()
		_, err = io.Copy(out, file)
		if err != nil {
			return "", rest_errors.NewInternalServerErr("Error while saving Pic", err)
		}
		return fileName, nil
	} else {
		return "", rest_errors.NewRestError("File Already exist", http.StatusAlreadyReported, "Already exist", nil)
	}
}

func UpdateFile(header *multipart.FileHeader, file multipart.File, path string) (string, rest_errors.RestErr) {
	splittedPath := strings.Split(path, "/")
	fileName := splittedPath[len(splittedPath)-1]
	DeleteFile(fileName)
	newFileName, err := SaveFile(header, file)
	if err != nil {
		return "", err
	}
	return newFileName, nil
}
