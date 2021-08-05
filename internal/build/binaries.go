package build

import (
	"os"
	"os/exec"
	"sync"

	"github.com/sudonims/hover/internal/log"
)

type binLookup struct {
	Name                string
	InstallInstructions string
	fullPath            string
	once                sync.Once
}

func (b *binLookup) FullPath() string {
	b.once.Do(func() {
		var err error
		b.fullPath, err = exec.LookPath(b.Name)
		if err != nil {
			log.Errorf("Failed to lookup `%s` executable: %s. %s", b.Name, err, b.InstallInstructions)
			os.Exit(1)
		}
	})
	return b.fullPath
}

var (
	goBinLookup = binLookup{
		Name:                "go",
		InstallInstructions: "Please install go or add `--docker` to run the Hover command in a Docker container.\nhttps://golang.org/doc/install",
	}
	flutterBinLookup = binLookup{
		Name:                "flutter",
		InstallInstructions: "Please install flutter or add `--docker` to run the Hover command in Docker container.\nhttps://flutter.dev/docs/get-started/install",
	}
	gitBinLookup = binLookup{
		Name: "git",
	}
	dockerBinLookup = binLookup{
		Name: "docker",
	}
)

func GoBin() string {
	return goBinLookup.FullPath()
}

func FlutterBin() string {
	return flutterBinLookup.FullPath()
}

func GitBin() string {
	return gitBinLookup.FullPath()
}

func DockerBin() string {
	return dockerBinLookup.FullPath()
}
