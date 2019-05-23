package scaffolding

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/config"
)

const (
	orionDirName     = "orion/"
	templatePathName = "infra/templates/"
)

// GitProvider represents how we will scaffold and uses git to provide
// the files to the workspace.
type GitProvider struct {
	Workspace Workspace
	Repo      Repo
}

// NewGitProvider ctor function for building a GitProvider struct
func NewGitProvider(repo Repo, workspace Workspace) GitProvider {
	gp := GitProvider{}
	gp.Repo = repo
	gp.Workspace = workspace

	return gp
}

// FetchFiles uses git to fetch files from the repo into the workspace.
func (g GitProvider) FetchFiles() (Workspace, func()) {
	remoteName := "origin"

	cleanup := g.Workspace.startTemporaryWorkspace()

	// git init
	r, err := git.PlainInit(g.Workspace.TemporaryLocationPath, false)
	if err != nil {
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
	err = workTree.Pull(&git.PullOptions{RemoteName: remoteName, Depth: 1})
	if err != nil {
		panic(err)
	}

	return g.Workspace, cleanup
}

// Workspace represents the workspace used for scoffolding. It
// contains both temporary and working directories
type Workspace struct {
	TemporaryLocationPath    string
	SourceDirectoryPath      string
	DestinationDirectoryPath string
}

// NewWorkspace ctor function for building a workspace struct
func NewWorkspace(templateParam string) Workspace {
	ws := Workspace{}

	_, templateName := parseParam(templateParam)

	ws.TemporaryLocationPath = func() string {
		filePathString := os.TempDir() + orionDirName
		filePath := filepath.FromSlash(filePathString)

		return filePath
	}()

	ws.SourceDirectoryPath = func(temporaryLocationPath string, templateName string) string {
		filePathString := temporaryLocationPath + templatePathName + templateName + "/"
		filePath := filepath.FromSlash(filePathString)

		return filePath
	}(ws.TemporaryLocationPath, templateName)

	ws.DestinationDirectoryPath = func() string {
		wd, err := os.Getwd()
		if err != nil {
			panic(err)
		}

		filePathString := wd + "/"
		filePath := filepath.FromSlash(filePathString)

		return filePath
	}()

	return ws
}

// startTemporaryWorkspace #longnamebro creates the temp
// dir that we will clone the repo into. Returns
// cleanup function that is intended for use later after work is complete.
func (w Workspace) startTemporaryWorkspace() func() {
	if _, err := os.Stat(w.TemporaryLocationPath); os.IsNotExist(err) {
		err := os.Mkdir(w.TemporaryLocationPath, 0700)
		if err != nil {
			panic("Could not create temporary workspace.")
		}
	}

	// Clean-up closure to be called after work in temporary directory
	// is completed.
	return func() {
		if _, err := os.Stat(w.TemporaryLocationPath); os.IsExist(err) {
			err := os.RemoveAll(w.TemporaryLocationPath)
			if err != nil {
				panic("Could not delete temorary workspace.")
			}
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
	repo := Repo{}

	repo.RepositoryName, repo.TemplateName = parseParam(templateParam)

	repo.RepositoryURL = func(repositoryName string) string {
		return fmt.Sprintf("https://github.com/microsoft/%s.git", repositoryName)
	}(repo.RepositoryName)

	return repo
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
