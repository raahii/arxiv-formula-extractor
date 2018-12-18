package utils

import (
	"io"
	"net/http"
	"os"
)

func downloadTarball(paperId string, path string) error {
	url := "https://arxiv.org/e-print/" + paperId

	// Create the file
	out, err := os.Create(path)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func main() {
}
