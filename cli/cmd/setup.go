package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	git "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/config"
)

var (
	template, repository, owner, source string
)

var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Setup a new project",
	Long:  `Setup command is used to create a project with templates`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(template) > 0 {
			fmt.Printf("orion setup --template is %s\n", template)

			allTheThings()
		}
	},
}

// TODO: should make these truly required/optional
func init() {
	rootCmd.AddCommand(setupCmd)
	setupCmd.Flags().StringVarP(&template, "template", "t", "", "Template name. Example: \"azure-simple\"")
	setupCmd.Flags().StringVarP(&repository, "repo", "r", "", "Repository name. Example: \"bedrock\"")
	setupCmd.Flags().StringVarP(&owner, "owner", "o", "", "Owner of repo's name. Example: \"Microsoft\"")
	setupCmd.Flags().StringVarP(&source, "source", "s", "", "Optional - Source Control host. Default: \"github.com\"")
}

func getWorkingDirectory() string {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	return wd
}

func getTempPath() string {
	tmpDir := os.TempDir()
	tmpDir += "\\orion"
	return tmpDir
}

// TODO: may not need to do this... could possibly use memfs?
func upsertTempDirectory() string {
	tmpDir := getTempPath()

	if _, err := os.Stat(tmpDir); os.IsNotExist(err) {
		err := os.Mkdir(tmpDir, 0700)
		if err != nil {
			panic(err)
		}
	}

	return tmpDir
}

func deleteTempDirectory() {
	tmpDir := getTempPath()

	if _, err := os.Stat(tmpDir); !os.IsNotExist(err) {
		err := os.RemoveAll(tmpDir)
		if err != nil {
			panic(err)
		}
	}
}

// TODO: replace params with struct to fix the optional call
func createGitRemoteURL(owner string, repo string, sourceOptional string) string {
	source := "github.com"

	if len(sourceOptional) > 0 {
		// TODO: should probably check the value and verify for correctness
		source = sourceOptional
	}

	url := fmt.Sprintf("https://%s/%s/%s.git", source, owner, repo)

	// TODO: should verify actual clone url
	return url
}

func allTheThings() {
	deleteTempDirectory()

	// TODO: move this to const file or something
	remoteName := "origin"

	// Make temp directory
	tmpPath := upsertTempDirectory()

	// git init
	r, err := git.PlainInit(tmpPath, false)
	if err != nil {
		panic(err)
	}

	// git remote add origin git@github.com:Microsoft/orion.git;
	remoteURL := createGitRemoteURL(owner, repository, source)
	_, err = r.CreateRemote(&config.RemoteConfig{
		Name: remoteName,
		URLs: []string{remoteURL},
	})
	if err != nil {
		panic(err)
	}

	// git fetch --depth 1;
	fetchOptions := &git.FetchOptions{
		RemoteName: remoteName,
		Depth:      1,
	}
	err = fetchOptions.Validate()
	if err != nil {
		panic(err)
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

	// This code is ref for when tree-ish checkout is available. Remove pull code above.
	// master := plumbing.ReferenceName("origin/master") // .NewBranchReferenceName("master")
	// err = workTree.Checkout(&git.CheckoutOptions{
	// 	Branch: master,
	// })
	// if err != nil {
	// 	panic(err)
	// }

	// Move module from tmp to working dir
	pwd := "C:\\tmp\\tmpgit\\infra" //getWorkingDirectory()
	infraPath := tmpPath + "\\infra"
	err = os.Rename(infraPath, pwd)
	if err != nil {
		panic(err)
	}

	deleteTempDirectory()
}
