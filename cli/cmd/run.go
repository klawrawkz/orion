package cmd

import (
	"log"
	"os"
	"os/exec"
	"runtime"

	"github.com/microsoft/orion/cli/pkg/download"
	"github.com/spf13/cobra"
)

var (
	docker                bool
	localRun              = "https://raw.githubusercontent.com/microsoft/cobalt/master/test-harness/local-run.sh"
	localRunWithOutDocker = "https://raw.githubusercontent.com/microsoft/cobalt/master/test-harness/local-run-wo-docker.sh"
	initScript            = "https://raw.githubusercontent.com/microsoft/cobalt/master/test-harness/init.sh"
	dockerfile            = "https://raw.githubusercontent.com/microsoft/cobalt/master/test-harness/Dockerfile"
	mageFile              = "https://raw.githubusercontent.com/microsoft/cobalt/master/test-harness/magefile.go"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run local test harness",
	Long:  `Run local test harness`,
	Run: func(cmd *cobra.Command, args []string) {
		if runtime.GOOS == "windows" {
			log.Fatalln("This will not run on Windows")
		}

		var testHarnessSetupScript download.URL
		initScriptURL := download.NewURL(initScript)
		dockerfileURL := download.NewURL(dockerfile)
		mageFileURL := download.NewURL(mageFile)

		if docker {
			testHarnessSetupScript = download.NewURL(localRun)
		} else {
			testHarnessSetupScript = download.NewURL(localRunWithOutDocker)
		}

		urls := []download.URL{
			initScriptURL,
			dockerfileURL,
			mageFileURL,
			testHarnessSetupScript,
		}

		dlManager := download.NewManager(urls)
		dlManager.FetchAll()

		runScript(testHarnessSetupScript.FileName)
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
	runCmd.Flags().BoolVarP(&docker, "docker", "d", true, "Run test harness in local docker")
}

func runScript(script string) {
	cmd := exec.Command("sh", script)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		log.Fatalf("Test harness failed to run with error: %s\n", err)
	}
}

//   |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |
//   |                                            |
//   |     NNNN        NN  II  NN        NNNN     |
//   |     NN  NN      NN  II  NN      NN  NN     |
//   |     NN    NN    NN  II  NN    NN    NN     |
//   |     NN      NN  NN  II  NN  NN      NN     |
//   |     NN        NNNN  II  NNNN        NN     |
//   |                                            |
//   |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |
