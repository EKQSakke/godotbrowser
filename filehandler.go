package main

import (
	"io"
	"net/http"
	"os"
)

func Download(url string, fileName string) error {
	filepath := "C:/temp/" + fileName
	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}