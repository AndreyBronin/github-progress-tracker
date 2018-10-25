/*
 *    Copyright 2018 Andrey Bronin <jonnib@yandex.ru>.
 *
 *    Licensed under the Apache License, Version 2.0 (the "License");
 *    you may not use this file except in compliance with the License.
 *    You may obtain a copy of the License at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 *    Unless required by applicable law or agreed to in writing, software
 *    distributed under the License is distributed on an "AS IS" BASIS,
 *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *    See the License for the specific language governing permissions and
 *    limitations under the License.
 */

/*
Package tracker provides functionality to track development progress in selected Github repos
*/
package tracker

import (
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

// NewAuthenticatedGithubTracker creates new tracker with auth
func NewAuthenticatedGithubTracker(storage Storage) (*GithubTracker, error) {
	// httpClient := oauth2.NewClient() // github oauth2
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

	// get data from github
}



// ProcessOwnerRepos collects data from all repos of particular owner
func (c *GithubTracker) ProcessOwnerRepos(owner string, repos []string) error {
	/*
	var err error
	if len(repos) == 0 {
		repos, err  = c.GetOwnerRepos(owner)
		if err != nil {
			return errors.Wrap(err, "Failed to GetOwnerRepos.")
		}
	}
*/
	// TODO: implement me

	return nil
}

