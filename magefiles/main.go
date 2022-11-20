//go:build mage

package main

import (
	"context"
	"os"
	"runtime"

	"dagger.io/dagger"
	gabadge "github.com/abtris/ga-badge/cli/build"
)

func GoGaBadge(ctx context.Context) {
	client, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stdout))
	if err != nil {
		panic(err)
	}
	defer client.Close()
	os, arch := getOsArch()

	uname := gabadge.Build(ctx, client, os, arch)

	_, err = uname.Export(ctx, ".")
	if err != nil {
		panic(err)
	}
}

func getOsArch() (string, string) {
	return runtime.GOOS, runtime.GOARCH
}
