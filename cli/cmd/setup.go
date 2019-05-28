package cmd

import (
	"github.com/microsoft/orion/cli/scaffolding"
	"github.com/spf13/cobra"
)

const (
	gitProvider    = true
	githubProvider = false
)

var template string
var provider scaffolding.Provider

var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Setup a new project from an existing template",
	Long:  `Setup command is used to create a project with templates`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(template) > 0 {
			repo := scaffolding.NewRepo(template)
			workspace := scaffolding.NewWorkspace(template)

			if gitProvider {
				provider = scaffolding.NewGitProvider(repo, workspace)
			}

			if githubProvider {
				// TODO: provider needed
			}

			s := scaffolding.NewScaffold(provider)
			s.Generate()
		} else {
			panic("must provide template param")
		}
	},
}

func init() {
	rootCmd.AddCommand(setupCmd)
	setupCmd.Flags().StringVarP(&template, "template", "t", "", "Template name. Example: \"cobalt/azure-simple\"")
}
