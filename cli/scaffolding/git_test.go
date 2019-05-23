package scaffolding

import (
	"testing"
)

func Test_parseParam(t *testing.T) {
	expectedRepo := "cobalt"
	expectedTemplate := "azure-simple"
	actualRepo, actualTemplate := parseParam("cobalt/azure-simple")

	if actualRepo != expectedRepo {
		t.Errorf("Expected repo %s to be %s", actualRepo, expectedRepo)
	}

	if actualTemplate != expectedTemplate {
		t.Errorf("Expected template %s to be %s", actualTemplate, expectedTemplate)
	}
}
