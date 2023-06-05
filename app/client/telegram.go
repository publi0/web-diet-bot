package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type fileInfo struct {
	Ok     bool `json:"ok"`
	Result struct {
		FileId       string `json:"file_id"`
		FileUniqueId string `json:"file_unique_id"`
		FileSize     int    `json:"file_size"`
		FilePath     string `json:"file_path"`
	} `json:"result"`
}

type TelegramClientAPI interface {
	DownloadImage(fileId string) ([]byte, error)
}

type TelegramClient struct {
	Client *http.Client
}

func (t TelegramClient) DownloadImage(fileId string) ([]byte, error) {
	token := os.Getenv("TELEGRAM_TOKEN")
	imageInfoURL := fmt.Sprintf("https://api.telegram.org/bot%s/getFile?file_id=%s", token, fileId)

	response, err := t.Client.Get(imageInfoURL)
	if err != nil {
		return nil, fmt.Errorf("error getting file path: %v", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error getting file path. Status: %v", response.StatusCode)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	var file fileInfo
	json.Unmarshal(body, &file)

	imageURL := fmt.Sprintf("https://api.telegram.org/file/bot%s/%s", token, file.Result.FilePath)
	response, err = t.Client.Get(imageURL)
	if err != nil {
		return nil, fmt.Errorf("error downloading image: %v", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error downloading image. Status: %v", response.StatusCode)
	}

	imageContent, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading image content: %v", err)
	}

	return imageContent, nil

}
