package scaffolding

import (
	"io/ioutil"
	"os"
)

// Provider interface encapsulates the behavior for fetching files
// from a repo.
type Provider interface {
	FetchFiles() *Workspace
}

// Scaffold uses a supplied provider to scaffold out a new
// or existing project
type Scaffold struct {
	Provider Provider
}

// NewScaffold ctor to build scaffold struct
func NewScaffold(provider Provider) Scaffold {
	return Scaffold{
		Provider: provider,
	}
}

// Generate performs the scaffolding
// should be GitRepo -> provider /Workspace/Scaffold - repo/workspace
func (s *Scaffold) Generate() {
	if s.Provider == nil {
		panic("Scaffold has no provider")
	}

	// Fetch files using provider
	workspace := s.Provider.FetchFiles()
	defer workspace.cleanupTemporaryWorkspace()

	// interate over all FS objects
	files, err := ioutil.ReadDir(workspace.SourceDirectoryPath)
	if err != nil {
		panic("Unable to read files from source directory")
	}

	// move files/folders to destination
	for _, file := range files {
		sourceFile := workspace.SourceDirectoryPath + file.Name()
		destinationFile := workspace.DestinationDirectoryPath + file.Name()

		err := os.Rename(sourceFile, destinationFile)
		if err != nil {
			print(err)
		}
	}

}
