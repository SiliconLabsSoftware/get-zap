package github

import (
	"context"
	"fmt"

	"github.com/google/go-github/github"
)

func SelfCheck() {
	client := github.NewClient(nil)

	fmt.Println("ALL REPOS:")
	opt := &github.RepositoryListByOrgOptions{Type: "public"}
	repos, _, err := client.Repositories.ListByOrg(context.Background(), "project-chip", opt)

	if err != nil {
		panic(err)
	}

	for _, repo := range repos {
		fmt.Println(*repo.Name)
	}

	fmt.Println("\nLATEST RELEASE:")
	release, _, err := client.Repositories.GetLatestRelease(context.Background(), "project-chip", "zap")
	if err != nil {
		panic(err)
	}
	fmt.Println(*release.TagName)
	//println(*release.AssetsURL)

	fmt.Println("\nALL RELEASES:")
	lo := &github.ListOptions{}
	releases, _, err := client.Repositories.ListReleases(context.Background(), "project-chip", "zap", lo)
	if err != nil {
		panic(err)
	}

	for _, release := range releases {
		fmt.Println(*release.TagName)
	}
}
