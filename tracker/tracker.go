/*
Package tracker provides functionality to track development progress in selected Github repos
 */
package tracker

import (
	"context"
	"fmt"
	"github.com/google/go-github/github"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
	"log"
	"os"

)

// Score is
type Score uint32

type Counters struct{
	commits uint32
	contributors uint32
	pullRequests uint32
}

// NewGithubTracker creates new tracker
func NewGithubTracker() (*GithubTracker, error) {
	return &GithubTracker{client: github.NewClient(nil)}, nil

}

type GithubTracker struct {
	client *github.Client
	organization string
	repos []string
}


func (c *GithubTracker) openRepo(organization, name string) (*git.Repository, error) {
	var repo *git.Repository
	path := fmt.Sprintf("./tmp/%s/%s", organization, name)
	repo, err := git.PlainOpen(path)
	if err == nil {
		repo.Fetch(&git.FetchOptions{Progress:os.Stdout, Force:true})

		//w, _ := repo.Worktree()
		//w.Pull(&git.PullOptions{Progress:os.Stdout, Force:true})
		return repo, nil
	}


	return git.PlainClone(path, true, &git.CloneOptions{
		URL: fmt.Sprintf("https://github.com/%s/%s", organization, name),
		Progress: os.Stdout,
	})
}

func (c *GithubTracker) ProcessRepo(organization, name string) {
	r, err := c.openRepo(organization, name)

	if err != nil {
		log.Fatalln(err.Error())
	}

	iter,_ := r.Log(&git.LogOptions{})

	var counter int
	iter.ForEach(func (c *object.Commit) error {
		counter++
		//fmt.Println(counter, c.Author.When, c.Author.Name, c.Hash)
		return nil
	})
	fmt.Println(counter)
}

func (c *GithubTracker)ProcessOrganizationRepos(organization string , names []string) error {
	for _, n := range names {
		for page := 0;;page++ {

			pullrequests, _, err := c.client.PullRequests.List(context.Background(), organization, n,
				&github.PullRequestListOptions{State: "closed", ListOptions: github.ListOptions{Page: page, PerPage: 100}})
			if err != nil {
				return err
			}

			for _, p := range pullrequests {
				fmt.Println( p.GetCreatedAt(), p.GetTitle())
			}

			if len(pullrequests) < 100 {
				break
			}
		}
	}

	return nil
}

