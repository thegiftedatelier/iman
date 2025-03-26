package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	imagePath := "path/to/your/image.jpg"
	boardID := "your_pinterest_board_id"
	accessToken := "your_pinterest_access_token"

	err := uploadImageToPinterest(imagePath, boardID, accessToken)
	if err != nil {
		fmt.Println("Error uploading image:", err)
	} else {
		fmt.Println("Image uploaded successfully!")
	}
}

func uploadImageToPinterest(imagePath, boardID, accessToken string) error {
	file, err := os.Open(imagePath)
	if err != nil {
		return err
	}
	defer file.Close()

	imageData, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	url := "https://api.pinterest.com/v1/pins/"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(imageData))
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", "Bearer "+accessToken)
	req.Header.Set("Content-Type", "multipart/form-data")

	q := req.URL.Query()
	q.Add("board", boardID)
	q.Add("note", "Uploaded via Iman Tool")
	req.URL.RawQuery = q.Encode()

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		body, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("failed to upload image: %s", body)
	}

	return nil
}
