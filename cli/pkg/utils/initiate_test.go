package utils

import (
	"testing"
	"os"
	"bytes"
	"io"

	"github.com/stretchr/testify/assert"
)

func Test_RunScript(t *testing.T) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	
	RunScript("../../scripts/echor.sh")
	
	outS := make(chan string)
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outS <- buf.String()
		}()
		
	w.Close()
	os.Stdout = old
	out := <-outS

	assert.Equal(t, "Ian was here\n", out)
}