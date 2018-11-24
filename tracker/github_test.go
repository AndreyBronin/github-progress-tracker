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
	"testing"
)

func TestGithubTracker_GetGithubContributors(t *testing.T) {
	t.Skip("api limits on travis")

	tracker, err := NewGithubTracker()
	assert.NoError(t, err)
	contributors, err := tracker.GetGithubContributors("insolar", "insolar")
	assert.NoError(t, err)
	assert.NotEqual(t, 0, len(contributors))
}

func TestGithubTracker_GetOwnerRepos(t *testing.T) {
	t.Skip("api limits on travis")

	tracker, err := NewGithubTracker()
	assert.NoError(t, err)

	repos, err := tracker.GetOwnerRepos("insolar")
	assert.NoError(t, err)
	assert.NotEqual(t, 0, len(repos))
}

func TestGithubTracker_GetPullRequests(t *testing.T) {
	t.Skip("api limits on travis")

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