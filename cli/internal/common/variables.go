package common

var (
	localRun              = gitContentsRoot + repo + branch + "test-harness/local-run.sh"
	localRunWithOutDocker = gitContentsRoot + repo + branch + "test-harness/local-run-wo-docker.sh"
	initScript            = gitContentsRoot + repo + branch + "test-harness/init.sh"
	dockerfile            = gitContentsRoot + repo + branch + "test-harness/Dockerfile"
	mageFile              = gitContentsRoot + repo + branch + "test-harness/magefile.go"
)

// GetTestHarnessFiles returns a string slice of the files needed to run the 
// test harness for Orion locally
func GetTestHarnessFiles() []string {
	return []string{
		localRun,
		localRunWithOutDocker,
		initScript,
		dockerfile,
		mageFile,
	}
}