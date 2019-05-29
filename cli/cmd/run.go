package cmd

import (
	"github.com/spf13/cobra"
)

var (
	docker bool
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run local test harness",
	Long:  `Run local test harness`,
	Run: func(cmd *cobra.Command, args []string) {
		remoteLocation := ""
		if docker {
			remoteLocation = getLocation()
		} else {
			remoteLocation = getLocationWithoutDocker()
		}
		localLocation, err := downloadScript(remoteLocation)
		if err != nil {
			panic(err)
		}
		runScript(localLocation)
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
	runCmd.Flags().BoolVarP(&docker, "docker", "d", true, "Run test harness in local docker")
}

func downloadInitSh() {

}

func downloadDockerfile() {

}

func downloadMageFile() {

}

func getLocation() string {
	return "https://raw.githubusercontent.com/microsoft/cobalt/master/test-harness/local-run.sh"
}

func getLocationWithoutDocker() string {
	return "https://raw.githubusercontent.com/microsoft/cobalt/master/test-harness/local-run-wo-docker.sh"
}

func downloadScript(scriptFile string) (string, error) {
	return "pwd", nil
}

func runScript(script string) {
	// Run script file
}
