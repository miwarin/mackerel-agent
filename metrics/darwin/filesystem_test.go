// +build darwin

package darwin

import (
	"testing"
)

func TestFilesystemGenerate(t *testing.T) {
	g := &FilesystemGenerator{}

	_, err := g.Generate()
	if err != nil {
		t.Errorf("Generate() failed: %s", err)
	}
}
