package arxiv

import (
	"fmt"
	"os"
	"testing"
)

func Exists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func Remove(path string) {
	if Exists(path) {
		if err := os.Remove(path); err != nil {
			fmt.Println(err)
		}
	}
}

func TestDownloadTarball(t *testing.T) {
	sampleUrl := "https://dev-files.blender.org/file/download/bwdp5reejwpkuh5i2oak/PHID-FILE-nui3bpuan4wdvd7yzjrs/sample.tar.gz"
	saveTo := "/tmp/sample.tar.gz"
	err := DownloadTarball(sampleUrl, saveTo)
	Remove(saveTo)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
}
