// Copyright © 2015 Steve Francia <spf@spf13.com>.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

/*
Package tracker provides functionality to track development progress in selected Github repos
*/
package tracker

import (
	"context"
	"fmt"
	"github.com/google/go-github/github"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"gopkg.in/src-d/go-git.v4"
	"os"
)

// Score is
type Score uint32

type Counters struct {
	commits      uint32
	contributors uint32
	pullRequests uint32
}

type GithubTracker struct {
	client       *github.Client
	organization string
	repos        []string
	storage      Storage
}

// NewGithubTracker creates new tracker
func NewGithubTracker() (*GithubTracker, error) {
	storage, err := NewStorage("commits_cache.sqlite3")
	if err != nil {
		return nil, errors.Wrap(err, "failed to init storage")
	}
	return &GithubTracker{client: github.NewClient(nil), storage: storage}, nil
}

// CloneRepo clones repo from Github opens cached repo
func (c *GithubTracker) CloneRepo(owner, name string) (*git.Repository, error) {
	var repo *git.Repository
	path := fmt.Sprintf("./github/%s/%s", owner, name)
	repo, err := git.PlainOpen(path)
	if err == nil {
		log.Infoln("Cache not found, cloning repo...")
		err = repo.Fetch(&git.FetchOptions{Progress: os.Stdout, Force: true})
		if err != nil {
			return nil, err
		}
		//w, _ := repo.Worktree()
		//w.Pull(&git.PullOptions{Progress:os.Stdout, Force:true})
		return repo, nil
	}

	return git.PlainClone(path, true, &git.CloneOptions{
		URL:      fmt.Sprintf("https://github.com/%s/%s", owner, name),
		Progress: os.Stdout,
	})
}

// ProcessRepo clones repo and collect all progress data
func (c *GithubTracker) ProcessRepo(name string, r *git.Repository) {

	iter, _ := r.Log(&git.LogOptions{})

	var counter int
	for {
		commit, err := iter.Next()
		if err != nil {
			break
		}

		counter++
		err = c.storage.SaveCommit(name, commit)
		if err != nil {
			log.Println(err.Error())
		}
	}

	log.Infof("Commits count: %d", counter)
}

// getOwnerRepos return all repos of the owner
func (c *GithubTracker) getOwnerRepos(owner string) ([]string, error) {
	var result []string
	for page := 0; ; page++ {

		repos, _, err := c.client.Repositories.List(context.Background(), owner,
			&github.RepositoryListOptions{
				Visibility: "public",
				Sort: "pushed",
				ListOptions: github.ListOptions{Page: page, PerPage: 100},
		})
		if err != nil {
			return nil, err
		}

		for _, r := range repos {
			result = append(result, r.GetName())
		}

		if len(repos) < 100 {
			break
		}
	}

	return result, nil
}

// ProcessOwnerRepos collects data from all repos of particular owner
func (c *GithubTracker) ProcessOwnerRepos(owner string, repos []string) error {
	for _, n := range repos {
		for page := 0; ; page++ {

			pullrequests, _, err := c.client.PullRequests.List(context.Background(), owner, n,
				&github.PullRequestListOptions{State: "closed", ListOptions: github.ListOptions{Page: page, PerPage: 100}})
			if err != nil {
				return err
			}

			for _, p := range pullrequests {
				fmt.Println(p.GetCreatedAt(), p.GetTitle())
			}

			if len(pullrequests) < 100 {
				break
			}
		}
	}

	return nil
}

func (c *GithubTracker) ProcessGithubContributors(organization, name string) error {
	ctx := context.Background()
	for page := 0; ; page++ {
		contributors, _, err := c.client.Repositories.ListContributors(ctx, organization, name, &github.ListContributorsOptions{
			ListOptions: github.ListOptions{Page: page, PerPage: 100},
		})

		if err != nil {
			return errors.Wrap(err, "Failed to get contributors list.")
		}

		for _, contributor := range contributors {
			fmt.Println(contributor.GetLogin(), contributor.GetType())

			if contributor.GetType() == "User" {
				//TODO: c.storage.SaveContributor()
			}
		}

		if len(contributors) < 100 {
			break
		}
	}

	return nil
}

/*
opt := &github.RepositoryListByOrgOptions{Type: "public"}
repos, _, err := client.Repositories.ListByOrg(context.Background(), "insolar", opt)
if err != nil {
	log.Fatalln(err.Error())
}

for _, r := range repos {
	fmt.Println(*r.Name)
}

events, _ , err := client.Activity.ListRepositoryEvents(context.Background(), "insolar", "insolar", &github.ListOptions{0, 500})
if err != nil {
	log.Fatalln(err.Error())
}

for _, e := range events {
	fmt.Println(e.GetCreatedAt(), e.GetType(), string(*e.RawPayload))
}
*/

//refs, _, _ := collector.client.Git.ListRefs(context.Background(), "insolar", "insolar", &github.ReferenceListOptions{ListOptions: github.ListOptions{0, 500}})
//for _, r := range refs {
//	fmt.Println(r.GetObject().GetType())
//}
/*
	err := collector.ProcessOwnerRepos("insolar", []string{"insolar"})
	if err != nil {
		log.Fatalln(err.Error())
	}
*/
//client.Repositories.ListContributors()
