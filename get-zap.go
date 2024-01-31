package main

import (
	"context"

	"github.com/google/go-github/v58/github"
)

func main() {
	client := github.NewClient(nil)

	opt := &github.RepositoryListByOrgOptions{Type: "public"}
	repos, _, err := client.Repositories.ListByOrg(context.Background(), "project-chip", opt)

	if err != nil {
		panic(err)
	}

	for _, repo := range repos {
		println(*repo.Name)
	}

	println("LATEST RELEASE:")
	release, _, err := client.Repositories.GetLatestRelease(context.Background(), "project-chip", "zap")
	if err != nil {
		panic(err)
	}
	println(*release.TagName)
	//println(*release.AssetsURL)

	println("ALL RELEASES:")
	lo := &github.ListOptions{}
	releases, _, err := client.Repositories.ListReleases(context.Background(), "project-chip", "zap", lo)
	if err != nil {
		panic(err)
	}

	for _, release := range releases {
		println(*release.TagName)
	}
}
