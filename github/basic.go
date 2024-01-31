package github

import (
	"context"
	"fmt"

	"github.com/google/go-github/github"
	"github.com/spf13/cobra"
)

func DownloadLatestRelease(owner string, repo string) {
	fmt.Printf("Downloading latest release for %v/%v...\n", owner, repo)
}

func SelfCheck(owner string, repo string) {
	client := github.NewClient(nil)

	// Get latest release
	fmt.Printf("Latest release of %v/%v:\n", owner, repo)
	release, _, err := client.Repositories.GetLatestRelease(context.Background(), owner, repo)
	cobra.CheckErr(err)
	fmt.Printf("  %v  [%v]\n", *release.TagName, *release.AssetsURL)

	// Get all releases
	fmt.Printf("\nAll releases of %v/%v:\n", owner, repo)
	lo := &github.ListOptions{}
	releases, _, err := client.Repositories.ListReleases(context.Background(), owner, repo, lo)
	cobra.CheckErr(err)
	for _, release := range releases {
		fmt.Printf("  %v  [%v]\n", release.GetTagName(), release.GetAssetsURL())
	}
}
