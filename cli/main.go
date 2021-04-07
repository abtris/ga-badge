package main

import (
	"embed"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"strings"

	"github.com/urfave/cli/v2"
)

// generateBadge
func generateBadge(githubActionURL string, branch string, label string) (string, error) {
	u, err := url.Parse(githubActionURL)
	if err != nil {
		log.Fatal(err)
	}
	parts := strings.Split(u.Path, "/")
	repoOwner := parts[1]
  repoName := parts[2]
  actionsString := parts[3]
	workflowsString := parts[4]
	if actionsString != "actions" || workflowsString != "workflows" {
		return "", fmt.Errorf("invalid URL on input")
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
    err := os.Mkdir(dirName, os.ModeDir)
    if err == nil || os.IsExist(err) {
        return nil
    } else {
        return err
    }
}

func createFile(content []byte, filename string) error {
	if _, err := os.Stat(".github/workflows/"+ filename); os.IsNotExist(err) {
  	err := ioutil.WriteFile(".github/workflows/"+ filename, content, os.ModePerm)
		if err != nil {
				return err
		}
	}
	return nil
}

func initWorkflow(lang string, templates []fs.DirEntry) (string, error) {
	err := ensureDir(".github/workflows")
	if err != nil {
		return "", err
	}
	fileName := fmt.Sprintf("%s.yaml", lang)
	switch lang {
	case "go":
		// "go.yaml"
		data, _ := files.ReadFile("templates/go.yaml")
		createFile(data, fileName)
	case "node","node.js":
		data, _ := files.ReadFile("templates/node.js.yaml")
		createFile(data, fileName)
	}
	return fileName, nil
}

//go:embed templates/*.yaml
var files embed.FS

func main()  {
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
				url := c.String("url");
				branch := c.String("branch");
				label := c.String("label");
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
		},{
			Name:    "init",
			Aliases: []string{"i"},
			Usage:   "Create github workflow file",
 			Flags: []cli.Flag{
        &cli.StringFlag{Name: "lang", Aliases: []string{"l"}, Value: "node.js", Required: true},
      },
			Action: func(c *cli.Context) error {
				lang := c.String("lang");
				fileName, err := initWorkflow(lang, templates)
				if err != nil {
					return err
				}
				fmt.Printf("Done. File .github/workflows/%s was created.\n", fileName)

				return nil
			},
		}}

  err := app.Run(os.Args)
  if err != nil {
    log.Fatal(err)
  }
}
