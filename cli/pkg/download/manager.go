package download

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

// URL contains the url and filename of the item to be downloaded
type URL struct {
	FileName string
	URL      string
}

// NewURL creates a URL from a url string
func NewURL(url string) URL {
	fileName := strings.Split(url, "/")

	return URL{
		FileName: fileName[len(fileName)-1],
		URL:      url,
	}
}

// Manager contains urls and functionality for downloading those items
type Manager struct {
	Urls []URL
}

// NewManager takes a slice of URLs and returns a Manager
func NewManager(urls []URL) Manager {
	return Manager{
		Urls: urls,
	}
}

// FetchAll will download all the files in the Manager.Urls slice
func (m *Manager) FetchAll() {
	ch := make(chan string)

	for i := range m.Urls {
		go fetch(m.Urls[i], ch)
	}

	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
}

func fetch(url URL, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url.URL)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	defer resp.Body.Close()

	out, err := os.Create(url.FileName)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	defer out.Close()

	w, err := io.Copy(out, resp.Body)
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, w, url)
}
