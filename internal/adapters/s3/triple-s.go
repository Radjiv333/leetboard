package s3

import (
	"fmt"
	"leetboard/internal/adapters/api"
	"log/slog"
	"net/http"
)

func UploadImages(characters []api.RickAndMortyCharacters, logger *slog.Logger) error {
	url := "http://triple-s:8081/characters"
	req, err := http.NewRequest(http.MethodPut, url, nil)
	if err != nil {
		logger.Error("Could not connect to s3: ", "error", err.Error())
		return err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logger.Error("Could not connect to s3: ", "error", err.Error())
		return err
	}
	defer resp.Body.Close()

	logger.Info("Checking response status", "response status", resp.Status)

	charactersNumber := len(characters)

	for i := 0; i < charactersNumber; i++ {
		fmt.Printf("№%d Loading \"%s's\" image...\n", i, characters[i].Name)
		imageResp, err := http.Get(characters[i].ImageURL)
		if err != nil {
			logger.Error("Could not GET image", "method", http.MethodGet, "error", err.Error())
			return err
		}

		if imageResp.StatusCode != http.StatusOK {
			imageResp.Body.Close()
			logger.Error("Could not GET image", "status", imageResp.Status)
			return fmt.Errorf("image GET failed: %s", imageResp.Status)
		}

		url := fmt.Sprintf("http://triple-s:8081/characters/%d", characters[i].ID)

		putReq, err := http.NewRequest(http.MethodPut, url, imageResp.Body)
		if err != nil {
			imageResp.Body.Close()
			logger.Error("Could not create PUT request to triple-s", "error", err.Error())
			return err
		}

		contentType := imageResp.Header.Get("Content-Type")
		if contentType != "" {
			putReq.Header.Set("Content-Type", contentType)
		}

		putResp, err := client.Do(putReq)
		imageResp.Body.Close()
		if err != nil {
			logger.Error("Could not send PUT request to triple-s", "error", err.Error())
			return err
		}

		if putResp.StatusCode < 200 || putResp.StatusCode >= 300 {
			putResp.Body.Close()
			logger.Error("triple-s returned bad status", "status", putResp.Status)
			return fmt.Errorf("triple-s PUT failed: %s", putResp.Status)
		}

		putResp.Body.Close()
	}
	return nil
}
