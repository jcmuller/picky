package configfile_test

import (
	"testing"

	"github.com/jcmuller/choosy/configfile"
)

func TestFilePath(t *testing.T) {
	expected := "/home/foobar/.config/choosy/config"

	actual, err := configfile.FilePath("/home/foobar")

	if err != nil {
		t.Errorf("not supposed to get error: %v", err)
	}

	if actual != expected {
		t.Errorf("File paths don't match:\n  expected: %+v\n    actual: %+v", expected, actual)
	}
}

func TestFileContentsCallsError(t *testing.T) {
	called := false
	var foo = func() { called = true }

	configfile.FileContents("/foobar", foo)

	if !called {
		t.Errorf("command didn't start")
	}
}
