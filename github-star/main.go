package main

import (
	"context"
	"log"
	"os"
	"strings"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

func main() {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_AUTH_TOKEN")},
	)
	client := github.NewClient(oauth2.NewClient(ctx, ts))

	for _, arg := range os.Args[1:] {
		if !strings.HasPrefix(arg, "github.com/") {
			log.Fatal("unknown repository")
		}
		userrepo := strings.Split(arg[11:], "/")
		if len(userrepo) != 2 {
			log.Fatal("unknown repository")
		}
		resp, err := client.Activity.Star(ctx, userrepo[0], userrepo[1])
		if err != nil {
			log.Fatal(err)
		}
		resp.Body.Close()
		if resp.StatusCode != 200 {
			log.Fatal(resp.Status)
		}
	}
}
