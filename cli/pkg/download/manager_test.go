package download

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	localRun                 = "https://raw.githubusercontent.com/microsoft/cobalt/master/test-harness/local-run.sh"
	localRunFileName         = "local-run.sh"
	localRunWoDocker         = "https://raw.githubusercontent.com/microsoft/cobalt/master/test-harness/local-run-wo-docker.sh"
	localRunWoDockerFileName = "local-run-wo-docker.sh"
)

func Test_NewURL(t *testing.T) {
	assert := assert.New(t)

	u := NewURL(localRun)

	assert.Equal(localRunFileName, u.FileName, "NewURL didn't parse the correct file name")
}

func Test_NewManager(t *testing.T) {
	assert := assert.New(t)

	urls := []string{localRun, localRunWoDocker}

	m := NewManager(urls)

	assert.Equal(2, len(m.Urls), "NewManager did create Manager object")
}

func Test_FetchAll(t *testing.T) {
	assert := assert.New(t)

	tempPath := func() string {
		filePathString := os.TempDir() + "/orion_test/"
		filePath := filepath.FromSlash(filePathString)

		return filePath
	}()

	tempFileName := func(fileName string) string {
		filePathString := tempPath + fileName
		filePath := filepath.FromSlash(filePathString)

		return filePath
	}

	// Just do this to create the manager
	urls := []string{localRun, localRunWoDocker}
	m := NewManager(urls)

	lr := NewURL(localRun)
	lr.FileName = tempFileName(lr.FileName)
	lrd := NewURL(localRunWoDocker)
	lrd.FileName = tempFileName(lrd.FileName)
	testUrls := []URL{lr, lrd}

	// swap the urls with testable url filenames in tmp dir
	m.Urls = testUrls

	if _, err := os.Stat(tempPath); os.IsNotExist(err) {
		err := os.Mkdir(tempPath, 0700)
		if err != nil {
			panic("Could not create temporary workspace.")
		}
	}

	m.FetchAll()
	files, err := ioutil.ReadDir(tempPath)
	if err != nil {
		panic("Unable to read files from directory")
	}

	assert.Equal(2, len(files))

	if _, err := os.Stat(tempPath); !os.IsNotExist(err) {
		err := os.RemoveAll(tempPath)
		if err != nil {
			panic("Could not delete temporary workspace.")
		}
	}
}
