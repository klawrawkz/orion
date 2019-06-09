package download

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

// Target contains the url and filename of the item to be downloaded
type Target struct {
	FileName string
	URL      string
}

// NewTarget creates a Target from a url string
func NewTarget(url string) Target {
	fileName := strings.Split(url, "/")

	return Target{
		FileName: fileName[len(fileName)-1],
		URL:      url,
	}
}

// Manager contains urls and functionality for downloading those items
type Manager struct {
	Urls []Target
}

// NewManager takes a slice of URLs and returns a Manager
func NewManager(urls []string) Manager {
	var URLs []Target

	for i := range urls {
		URLs = append(URLs, NewTarget(urls[i]))
	}

	return Manager{
		Urls: URLs,
	}
}

// FetchAll will download all the files in the Manager.Urls slice
func (m *Manager) FetchAll() {
	for i := range m.Urls {
		result, err := fetch(m.Urls[i])
		if err != nil {
			log.Fatalln(err)
		}

		log.Println(result)
	}
}

func fetch(url Target) (string, error) {
	start := time.Now()
	resp, err := http.Get(url.URL)
	if err != nil {
		return "", fmt.Errorf("While getting %s: %v", url.URL, err)
	}
	defer resp.Body.Close()

	out, err := os.Create(url.FileName)
	if err != nil {
		return "", fmt.Errorf("While creating file %s: %v", url.FileName, err)
	}
	defer out.Close()

	w, err := io.Copy(out, resp.Body)
	if err != nil {
		return "", fmt.Errorf("while reading %s: %v", url, err)
	}

	secs := time.Since(start).Seconds()
	return fmt.Sprintf("%.2fs  %7d  %s", secs, w, url), nil
}
