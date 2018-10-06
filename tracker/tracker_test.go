package tracker

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
	"io/ioutil"
	"os"
	"testing"
	"time"
)

func createTestRepo(t *testing.T, name string) (*git.Repository, string, error) {
	dir, err := ioutil.TempDir("", "git")
	assert.NoError(t, err)

	repo, err := git.PlainInit(dir, false)
	assert.NoError(t, err)

	wt, err := repo.Worktree()
	assert.NoError(t, err)

	_, err = wt.Filesystem.Create(".keep")
	assert.NoError(t, err)

	hash, err := wt.Commit("initial commit", &git.CommitOptions{Author: &object.Signature{
		Name:  "Andrey",
		Email: "info@example.com",
		When:  time.Now()},
	})
	assert.NoError(t, err)

	fmt.Println("Commit: ", hash.String())
	return repo, dir, err
}

func TestGithubCollector_ProcessRepo(t *testing.T) {
	//t.Skip("tmp")
	tracker, err := NewGithubTracker()
	assert.NoError(t, err)

	repo, tmpRepoDir, err := createTestRepo(t, "testrepo")
	defer os.RemoveAll(tmpRepoDir)

	//repo, err := tracker.CloneRepo("andreybronin", "github-progress-tracker")
	assert.NoError(t, err)

	tracker.ProcessRepo("testrepo", repo)

}

func TestGithubTracker_ProcessGithubContributors(t *testing.T) {
	tracker, err := NewGithubTracker()
	assert.NoError(t, err)
	tracker.ProcessGithubContributors("andreybronin", "github-progress-tracker")
}
