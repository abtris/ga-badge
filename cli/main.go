package main

import (
	"embed"
	"errors"
	"fmt"
	"io/fs"

	"log"
	"net/url"
	"os"
	"strings"

	"github.com/urfave/cli/v2"
)

// All of this added by goreleaser using ldflags
var (
	// Short Git Commit Hash
	CommitHash string
	// Version vx.x.x
	Version string
	// Date of build
	BuildDate string
)

func printVersion() {
	fmt.Printf("gab version: %s-%s (%s)\n", Version, CommitHash, BuildDate[:10])
}

var errWrongURL = errors.New("wrong URL, can't generate badge")

// generateBadge
func generateBadge(githubActionURL string, branch string, label string) (string, error) {
	u, err := url.Parse(githubActionURL)
	if err != nil {
		log.Fatal(err)
		return "", errWrongURL
	}
	parts := strings.Split(u.Path, "/")
	if len(parts) <= 5 {
		return "", errWrongURL
	}
	repoOwner := parts[1]
	repoName := parts[2]
	actionsString := parts[3]
	workflowsString := parts[4]
	if actionsString != "actions" || workflowsString != "workflows" {
		return "", errWrongURL
	}
	workflowFileName := parts[5]
	defaultBranch := ""
	if branch != "master" {
		defaultBranch = fmt.Sprintf("?branch=%s", branch)
	}
	title := label
	return fmt.Sprintf("[![%s](https://github.com/%s/%s/actions/workflows/%s/badge.svg%s)](https://github.com/%s/%s/actions)", title, repoOwner, repoName, workflowFileName, defaultBranch, repoOwner, repoName), nil
}

func ensureDir(dirName string) error {
	if _, err := os.Stat(dirName); os.IsNotExist(err) {
		err := os.MkdirAll(dirName, os.ModePerm)
		return err
	}
	return nil
}

func createFile(content []byte, filename string) error {
	if _, err := os.Stat(".github/workflows/" + filename); os.IsNotExist(err) {
		err := os.WriteFile(".github/workflows/"+filename, content, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}

func initWorkflow(lang string, actionsDir string, templates []fs.DirEntry) (string, error) {
	err := ensureDir(actionsDir)
	if err != nil {
		return "", err
	}
	fileName := fmt.Sprintf("%s.yaml", lang)
	switch lang {
	case "go":
		// "go.yaml"
		data, _ := files.ReadFile("templates/go.yaml")
		createFile(data, fileName)
	case "node", "node.js":
		data, _ := files.ReadFile("templates/node.js.yaml")
		createFile(data, fileName)
	}
	return fileName, nil
}

//go:embed templates/*.yaml
var files embed.FS

func main() {
	templates, _ := fs.ReadDir(files, "templates")
	app := cli.NewApp()
	app.EnableBashCompletion = true
	app.Commands = []*cli.Command{
		{
			Name:    "create",
			Aliases: []string{"c"},
			Usage:   "Create github action badge from URL",
			Flags: []cli.Flag{
				&cli.StringFlag{Name: "url", Aliases: []string{"u"}},
				&cli.StringFlag{Name: "branch", Aliases: []string{"b"}, Value: "master"},
				&cli.StringFlag{Name: "label", Aliases: []string{"l"}, Value: "Build Status"},
			},
			Action: func(c *cli.Context) error {
				url := c.String("url")
				branch := c.String("branch")
				label := c.String("label")
				if c.NArg() > 0 {
					url = c.Args().Get(0)
				}
				badge, err := generateBadge(url, branch, label)
				if err != nil {
					return err
				}
				fmt.Println(badge)

				return nil
			},
		}, {
			Name:    "init",
			Aliases: []string{"i"},
			Usage:   "Create github workflow file",
			Flags: []cli.Flag{
				&cli.StringFlag{Name: "lang", Aliases: []string{"l"}, Value: "node.js", Required: true},
			},
			Action: func(c *cli.Context) error {
				lang := c.String("lang")
				fileName, err := initWorkflow(lang, ".github/workflows/", templates)
				if err != nil {
					return err
				}
				fmt.Printf("Done. File .github/workflows/%s was created.\n", fileName)

				return nil
			},
		}, {
			Name:    "version",
			Aliases: []string{"v"},
			Usage:   "Print version",
			Action: func(c *cli.Context) error {
				printVersion()
				return nil
			},
		}}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
