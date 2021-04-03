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
func generateBadge(githubActionURL string) (string, error) {
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
  workflowFileName := parts[4]
	title := "Build Status"
	return fmt.Sprintf("[![%s](https://github.com/%s/%s/actions/workflows/%s/badge.svg)](https://github.com/%s/%s/actions)", title, repoOwner, repoName, workflowFileName, repoOwner, repoName), nil
}

func main()  {
	app := cli.NewApp()
	app.EnableBashCompletion = true
	app.Commands = []*cli.Command{
		{
			Name:    "badge",
			Aliases: []string{"b"},
			Usage:   "create badge",
			Action: func(c *cli.Context) error {
				badge, err := generateBadge(c.Args().First())
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

