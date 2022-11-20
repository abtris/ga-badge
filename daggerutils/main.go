package daggerutils

import (
	"context"
	"path"

	"dagger.io/dagger"
)

const (
	golangImage = "golang:latest"
)

type GoBuildInput struct {
	Directory *dagger.Directory
	Client    *dagger.Client
	Ctx       context.Context
	Os        string
	Arch      string
	Workdir   string
}

func GoBuild(cfg GoBuildInput) *dagger.Directory {
	// Load image
	builder := cfg.Client.Container().From(golangImage)

	workdir := path.Join("/src", cfg.Workdir)

	builder = builder.WithMountedDirectory("/src", cfg.Directory).
		WithWorkdir(workdir).
		WithEnvVariable("GOARCH", cfg.Arch).
		WithEnvVariable("GOOS", cfg.Os)

	// Execute Command
	builder = builder.Exec(dagger.ContainerExecOpts{
		Args: []string{"go", "build", "-o", "ga-badge"},
	})

	return builder.Directory(workdir)
}
