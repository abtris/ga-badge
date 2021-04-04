package main

import (
	"fmt"
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

func main()  {
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
		},
	}

  err := app.Run(os.Args)
  if err != nil {
    log.Fatal(err)
  }
}

