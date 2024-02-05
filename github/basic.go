/*
Copyright © 2024 Silicon Labs
*/
package github

import (
	"context"
	"fmt"
	"io"

	"github.com/google/go-github/github"
	"github.com/spf13/cobra"
	"golang.org/x/oauth2"
)

type GithubConfiguration struct {
	Owner   string
	Repo    string
	Release string
	Token   string
}

func CreateGithubClient(cfg *GithubConfiguration) *github.Client {
	var client *github.Client
	if cfg.Token == "" {
		fmt.Println("You do not have GET_ZAP_TOKEN set. This will limit the number of requests you can make to the github API.")
		client = github.NewClient(nil)
	} else {
		ctx := context.Background()
		ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: cfg.Token})
		tc := oauth2.NewClient(ctx, ts)
		client = github.NewClient(tc)
	}
	return client
}

func DownloadRelease(cfg *GithubConfiguration) {

	client := CreateGithubClient(cfg)
	// Get latest release
	fmt.Printf("Latest release of %v/%v:\n", cfg.Owner, cfg.Repo)
	release, _, err := client.Repositories.GetLatestRelease(context.Background(), cfg.Owner, cfg.Repo)
	cobra.CheckErr(err)
	fmt.Printf("  %v  [%v]\n", *release.TagName, *release.AssetsURL)
	assets, _, err := client.Repositories.ListReleaseAssets(context.Background(), cfg.Owner, cfg.Repo, release.GetID(), &github.ListOptions{})
	cobra.CheckErr(err)
	for _, asset := range assets {
		fmt.Printf(" - Download %v [%v]\n", asset.GetName(), asset.GetSize())
		rc, redirect, err := client.Repositories.DownloadReleaseAsset(context.Background(), cfg.Owner, cfg.Repo, asset.GetID())
		cobra.CheckErr(err)
		if rc != nil {
			fmt.Printf("Not redirected.\n")
			buff := make([]byte, 10*1024)
			n, err := rc.Read(buff)
			fmt.Printf("Read %d bytes\n", n)
			if n > 0 {
				fmt.Printf("Read %d bytes\n", n)
			} else {
				if err == io.EOF {
					rc.Close()
					fmt.Println("Eof")
				} else {
					cobra.CheckErr(err)
				}
			}
		} else {
			DownloadFileFromUrl(redirect, asset.GetName(), DefaultSecurityOptions())
		}
	}
}

func printRelease(client *github.Client, owner string, repo string, release *github.RepositoryRelease) {
	fmt.Printf("  %v  [Published: %v]\n", release.GetTagName(), release.GetCreatedAt())
	assets, _, err := client.Repositories.ListReleaseAssets(context.Background(), owner, repo, release.GetID(), &github.ListOptions{})
	cobra.CheckErr(err)
	for _, asset := range assets {
		fmt.Printf("    %v [%v bytes]\n", asset.GetName(), asset.GetSize())
	}
}

func ListGithub(cfg *GithubConfiguration) {
	client := CreateGithubClient(cfg)
	if cfg.Release == "" {
		fmt.Printf("Listing all releases of repo '%v/%v':\n", cfg.Owner, cfg.Repo)
		allReleases, _, err := client.Repositories.ListReleases(context.Background(), cfg.Owner, cfg.Repo, &github.ListOptions{})
		cobra.CheckErr(err)
		for _, release := range allReleases {
			fmt.Printf("  %v [Published: %v]\n", release.GetTagName(), release.GetPublishedAt())
		}
	} else if cfg.Release == "latest" {
		// Get latest release
		fmt.Printf("Checking latest release of repo '%v/%v':\n", cfg.Owner, cfg.Repo)
		release, _, err := client.Repositories.GetLatestRelease(context.Background(), cfg.Owner, cfg.Repo)
		cobra.CheckErr(err)
		printRelease(client, cfg.Owner, cfg.Repo, release)
	} else {
		// Get specific release
		fmt.Printf("Checking latest release of repo '%v/%v':\n", cfg.Owner, cfg.Repo)
		allReleases, _, err := client.Repositories.ListReleases(context.Background(), cfg.Owner, cfg.Repo, &github.ListOptions{})
		cobra.CheckErr(err)
		wasFound := false
		for _, release := range allReleases {
			if release.GetTagName() == cfg.Release {
				release, _, err := client.Repositories.GetRelease(context.Background(), cfg.Owner, cfg.Repo, release.GetID())
				cobra.CheckErr(err)
				printRelease(client, cfg.Owner, cfg.Repo, release)
				wasFound = true
				break
			}
		}
		if !wasFound {
			fmt.Printf("Could not find a release with tag '%v'\n", cfg.Release)
		}
	}

}
