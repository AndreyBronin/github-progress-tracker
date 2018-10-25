package tracker

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGithubTracker_GetGithubContributors(t *testing.T) {
	tracker, err := NewGithubTracker()
	assert.NoError(t, err)
	contributors, err := tracker.GetGithubContributors("insolar", "insolar")
	assert.NoError(t, err)
	assert.NotEqual(t, 0, len(contributors))
}

func TestGithubTracker_GetOwnerRepos(t *testing.T) {
	tracker, err := NewGithubTracker()
	assert.NoError(t, err)

	repos, err := tracker.GetOwnerRepos("insolar")
	assert.NoError(t, err)
	assert.NotEqual(t, 0, len(repos))
}

func TestGithubTracker_GetPullRequests(t *testing.T) {
	tracker, err := NewGithubTracker()
	assert.NoError(t, err)

	pr, err := tracker.GetPullRequests("insolar", "insolar")
	assert.NoError(t, err)
	assert.NotEqual(t, 0, len(pr))

}

func TestGithubTracker_GetAPILimits(t *testing.T) {
	tracker, err := NewGithubTracker()
	assert.NoError(t, err)

	limits, err := tracker.GetAPILimits()
	assert.NoError(t, err)
	fmt.Println(limits.String())
}