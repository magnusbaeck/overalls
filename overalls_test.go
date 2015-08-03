package main

import (
	"io/ioutil"
	"os/exec"
	"testing"

	. "gopkg.in/bluesuncorp/assert.v1"
)

// NOTES:
// - Run "go test" to run tests
// - Run "gocov test | gocov report" to report on test converage by file
// - Run "gocov test | gocov annotate -" to report on all code and functions, those ,marked with "MISS" were never called
//
// or
//
// -- may be a good idea to change to output path to somewherelike /tmp
// go test -coverprofile cover.out && go tool cover -html=cover.out -o cover.html
//

func TestOveralls(t *testing.T) {

	args := []string{"-project=github.com/joeybloggs/overalls/test-files", "-covermode=count", "-debug"}

	cmd := exec.Command("overalls", args...)
	err := cmd.Run()
	Equal(t, err, nil)

	fileBytes, err := ioutil.ReadFile(srcPath + "github.com/joeybloggs/overalls/test-files/overalls.coverprofile")
	Equal(t, err, nil)
	Equal(t, len(fileBytes), 151)
}