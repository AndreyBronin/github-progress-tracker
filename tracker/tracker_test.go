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

