package cmd

import (
	"log"
	"runtime"

	"github.com/microsoft/orion/cli/pkg/download"
	"github.com/microsoft/orion/cli/pkg/utils"
	"github.com/microsoft/orion/cli/internal/common"
	"github.com/spf13/cobra"
)

var (
	docker bool
)

var runCmd = &cobra.Command{
	Use:   "test",
	Short: "Run local test harness",
	Long:  `Run local test harness`,
	Run: func(cmd *cobra.Command, args []string) {
		if runtime.GOOS == "windows" {
			log.Fatalln("This will not run on Windows")
		}

		urls := common.GetTestHarnessFiles()

		if docker == true {
			// Cut 2nd item in slice (without docker script)
			urls = append(urls[:1], urls[2:]...)
		} else {
			// Cut 1st item in slice (docker script)
			urls = urls[1:]
		}

		dlManager := download.NewManager(urls)
		dlManager.FetchAll()

		utils.RunScript("bash", dlManager.Urls[0].FileName)
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
	runCmd.Flags().BoolVarP(&docker, "docker", "d", false, "Run test harness in local docker")
}
