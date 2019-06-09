package scaffolding

import (
	"io/ioutil"
	"log"
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
	defer deleteFolderIfExist(workspace.TemporaryDirectoryPath)

	// Check if folder structure exists and create if doesn't
	checkIfRequestedTemplateExists(workspace)
	checkIfModulesExist(workspace)
	checkIfRootFolderExists(workspace)

	// mv requested template into template folder
	moveFiles(workspace.TemporaryTemplateDirectoryPath, workspace.DestinationTemplateDirectoryPath)

	// mv all modules/providers into folder...
	// TODO: should we figureout the provider?
	moveFiles(workspace.TemporaryModuleDirectoryPath, workspace.DestinationModuleDirectoryPath)

}

func checkIfRequestedTemplateExists(workspace *Workspace) {
	requestedTemplateExist := checkFolderExist(workspace.DestinationTemplateDirectoryPath)
	if requestedTemplateExist {
		log.Fatalln("Requested template already exists in current working directory.")
	} else {
		createFolderIfNotExist(rootPathName)
		createFolderIfNotExist(templatePathName)
		createFolderIfNotExist(workspace.DestinationTemplateDirectoryPath)
	}
}

func checkIfModulesExist(workspace *Workspace) {
	// Does Module folder already exist?
	// Yes, error message...
	// TODO: prompt for overwrite
	// Create Module folder
	modulesExist := checkFolderExist(modulePathName)

	if modulesExist {
		msg := "Module dir already exists. Please backup/remove this directory before proceeding."
		log.Fatalln(msg)
	} else {
		createFolderIfNotExist(modulePathName)
	}
}

func checkIfRootFolderExists(workspace *Workspace) {
	// Does Infra folder exist, yes log msg
	// Create Infra folder
	rootExist := checkFolderExist(rootPathName)

	if rootExist {
		log.Println("Infra/ folder already exists, skipping create.")
	} else {
		createFolderIfNotExist(rootPathName)
	}
}

func moveFiles(source string, destination string) {
	// interate over all FS objects
	files, err := ioutil.ReadDir(source)
	if err != nil {
		panic("Unable to read files from source directory")
	}

	// move files/folders to destination
	for _, file := range files {
		sourceFile := source + file.Name()
		destinationFile := destination + file.Name()

		err := os.Rename(sourceFile, destinationFile)
		if err != nil {
			print(err)
		}
	}
}
