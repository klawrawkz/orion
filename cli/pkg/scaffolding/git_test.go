package scaffolding

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
	"fmt"

	"github.com/stretchr/testify/assert"
)

const (
	templateParam = "cobalt/azure-simple"
	repo          = "cobalt"
	template      = "azure-simple"
)

func Test_NewGitProvider(t *testing.T) {
	gp := NewGitProvider(NewRepo(templateParam), NewWorkspace(templateParam))

	assert.NotNil(t, gp)
}

func Test_NewWorkspace(t *testing.T) {
	w := NewWorkspace(templateParam)

	assert.NotNil(t, w)
}

func Test_NewRepo(t *testing.T) {
	r := NewRepo(templateParam)

	assert.NotNil(t, r)
}

func Test_parseParam(t *testing.T) {
	expectedRepo := repo
	expectedTemplate := template
	actualRepo, actualTemplate := parseParam(templateParam)

	assert.Equal(t, actualRepo, expectedRepo, "Actual and Expected Repo should be equal")
	assert.Equal(t, actualTemplate, expectedTemplate, "Actual and Expected Template should be equal")
}

func Test_FetchFiles(t *testing.T) {
	assert := assert.New(t)

	gp := NewGitProvider(NewRepo(templateParam), NewWorkspace(templateParam))
	fmt.Println(gp.Workspace.TemporaryLocationPath)
	ws := gp.FetchFiles()
	defer ws.cleanupTemporaryWorkspace()

	_, err := ioutil.ReadDir(gp.Workspace.SourceDirectoryPath)
	assert.NoError(err)

}

func Test_startTemporaryWorkspace(t *testing.T) {
	assert := assert.New(t)

	w := NewWorkspace(templateParam)
	w.startTemporaryWorkspace()
	defer w.cleanupTemporaryWorkspace()

	expectedFilePath := filepath.FromSlash(os.TempDir() + orionDirName)

	if _, err := os.Stat(expectedFilePath); os.IsNotExist(err) {
		assert.Error(err)
	}
}
