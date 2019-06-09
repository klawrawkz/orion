package scaffolding

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/config"
	"gopkg.in/src-d/go-git.v4/plumbing"
)

const (
	orionDirName     = "orion/"
	rootPathName     = "infra/"
	templatePathName = rootPathName + "templates/"
	modulePathName   = rootPathName + "modules/"
)

// GitProvider represents how we will scaffold and uses git to provide
// the files to the workspace.
type GitProvider struct {
	Workspace Workspace
	Repo      Repo
}

// NewGitProvider ctor function for building a GitProvider struct
func NewGitProvider(repo Repo, workspace Workspace) GitProvider {
	return GitProvider{
		Repo:      repo,
		Workspace: workspace,
	}
}

// FetchFiles uses git to fetch files from the repo into the workspace.
func (g GitProvider) FetchFiles() *Workspace {
	remoteName := "origin"
	branch := "master"

	createFolderIfNotExist(g.Workspace.TemporaryDirectoryPath)

	// git init
	r, err := git.PlainInit(g.Workspace.TemporaryDirectoryPath, false)
	if err != nil {
		deleteFolderIfExist(g.Workspace.TemporaryDirectoryPath)
		panic(err)
	}

	// git remote add origin git@github.com:Microsoft/orion.git;
	_, err = r.CreateRemote(&config.RemoteConfig{
		Name: remoteName,
		URLs: []string{g.Repo.RepositoryURL},
	})
	if err != nil {
		panic(err)
	}

	// git fetch --depth 1;
	fetchOptions := &git.FetchOptions{
		RemoteName: remoteName,
		Depth:      1,
	}

	err = r.Fetch(fetchOptions)
	if err != nil {
		panic(err)
	}

	// git checkout origin/master infra ?or infra/templates/azure-simple-hw;
	workTree, err := r.Worktree()
	if err != nil {
		panic(err)
	}

	// We don't clone / pull because this is a stop gap until go-git implements
	// tree-ish checkout ability. This will enable us to only pull the needed
	// folder and not entire code base.
	// Branch logic here is for future use for pinning to tag/branch vs master.
	if branch == "master" {
		err = workTree.Pull(&git.PullOptions{RemoteName: remoteName, Depth: 1})
		if err != nil {
			panic(err)
		}
	} else {
		err = workTree.Pull(&git.PullOptions{
			RemoteName:    remoteName,
			Depth:         1,
			ReferenceName: plumbing.ReferenceName(branch),
		})
		if err != nil {
			panic(err)
		}
	}

	return &g.Workspace
}

// Workspace represents the workspace used for scoffolding. It
// contains both temporary and working directories
type Workspace struct {
	TemporaryDirectoryPath           string
	TemporaryTemplateDirectoryPath   string
	TemporaryModuleDirectoryPath     string
	DestinationDirectoryPath         string
	DestinationTemplateDirectoryPath string
	DestinationModuleDirectoryPath   string
}

// NewWorkspace ctor function for building a workspace struct
func NewWorkspace(templateParam string) Workspace {

	_, templateName := parseParam(templateParam)

	temporaryPath := func() string {
		filePathString := os.TempDir() + orionDirName
		filePath := filepath.FromSlash(filePathString)

		return filePath
	}()

	destinationPath := func() string {
		wd, err := os.Getwd()
		if err != nil {
			panic(err)
		}

		filePathString := wd + "/"
		filePath := filepath.FromSlash(filePathString)

		return filePath
	}()

	templatePath := func(basePath string, templateName string) string {
		filePathString := basePath + templatePathName + templateName + "/"
		filePath := filepath.FromSlash(filePathString)

		return filePath
	}

	modulePath := func(basePath string) string {
		filePathString := basePath + modulePathName
		filePath := filepath.FromSlash(filePathString)

		return filePath
	}

	return Workspace{
		TemporaryDirectoryPath:           temporaryPath,
		TemporaryTemplateDirectoryPath:   templatePath(temporaryPath, templateName),
		TemporaryModuleDirectoryPath:     modulePath(temporaryPath),
		DestinationDirectoryPath:         destinationPath,
		DestinationTemplateDirectoryPath: templatePath(destinationPath, templateName),
		DestinationModuleDirectoryPath:   modulePath(destinationPath),
	}
}

// createFolderIfNotExist will create a dir given a path if it doesn't
// already exist.
func createFolderIfNotExist(folderPath string) {
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		err := os.Mkdir(folderPath, 0700)
		if err != nil {
			panic(fmt.Sprintf("Could not create %s", folderPath))
		}
	}
}

func checkFolderExist(folderPath string) bool {
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		return false
	}

	return true
}

// deleteFolderIfExist removes a folder given a path if it exists.
func deleteFolderIfExist(folderPath string) {
	if _, err := os.Stat(folderPath); !os.IsNotExist(err) {
		err := os.RemoveAll(folderPath)
		if err != nil {
			panic(fmt.Sprintf("Could not delete %s", folderPath))
		}
	}
}

// Repo represents a repo that contains the template we want
type Repo struct {
	RepositoryName string
	RepositoryURL  string
	TemplateName   string
}

// NewRepo ctor function for building a repo struct
func NewRepo(templateParam string) Repo {

	repositoryName, templateName := parseParam(templateParam)

	repositoryURL := func(repositoryName string) string {
		return fmt.Sprintf("https://github.com/microsoft/%s.git", repositoryName)
	}(repositoryName)

	return Repo{
		RepositoryName: repositoryName,
		TemplateName:   templateName,
		RepositoryURL:  repositoryURL,
	}
}

// parseParam is a utility function that takes the template
// argument pased to orion setup cmd and parses and returns
// the repoName and templateName.
func parseParam(templateParam string) (string, string) {
	parsed := strings.Split(templateParam, "/")

	if len(parsed) != 2 {
		panic("The supplied Template parameter should be \"repo/template\"")
	}

	repositoryName := parsed[0]
	templateName := parsed[1]

	return repositoryName, templateName
}
