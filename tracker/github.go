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
	"context"
	"fmt"
	"github.com/google/go-github/v19/github"
	"github.com/pkg/errors"
)

func (c *GithubTracker) GetAPILimits() (*github.RateLimits, error) {
	limit, _ , err := c.client.RateLimits(context.Background())
	return limit, err
}

// GetOwnerRepos return all repos of the owner
func (c *GithubTracker) GetOwnerRepos(owner string) ([]*github.Repository, error) {
	result := make([]*github.Repository, 0)
	for page := 0; ; page++ {

		repos, _, err := c.client.Repositories.ListByOrg(context.Background(), owner,
			&github.RepositoryListByOrgOptions{
				Type: "public",
				ListOptions: github.ListOptions{Page: page, PerPage: 100},
			})
		if err != nil {
			return nil, err
		}

		result = append(result, repos...)
		if len(repos) < 100 {
			break
		}
	}

	return result, nil
}

// GetGithubContributors return all Contributors of the repo
func (c *GithubTracker) GetGithubContributors(owner, repo string) ([]*github.Contributor, error) {
	ctx := context.Background()

	result := make([]*github.Contributor, 0)
	for page := 0; ; page++ {
		contributors, _, err := c.client.Repositories.ListContributors(ctx, owner, repo, &github.ListContributorsOptions{
			ListOptions: github.ListOptions{Page: page, PerPage: 100},
		})

		if err != nil {
			return nil, errors.Wrap(err, "Failed to get contributors list.")
		}

		for _, contributor := range contributors {
			fmt.Println(contributor.GetLogin(), contributor.GetType())

			if contributor.GetType() == "User" {
				result = append(result, contributor)
			}
		}

		if len(contributors) < 100 {
			break
		}
	}

	return result, nil
}

// GetPullRequests returns all pull requests of the repo
func (c *GithubTracker) GetPullRequests(owner, repo string) ([]*github.PullRequest, error) {
	result := make([]*github.PullRequest, 0)
	for page := 0; ; page++ {
		pullrequests, _, err := c.client.PullRequests.List(context.Background(), owner, repo,
			&github.PullRequestListOptions{State: "closed", Sort: "created", ListOptions: github.ListOptions{Page: page, PerPage: 100}})
		if err != nil {
			return nil, errors.Wrap(err, "Failed to get pull requests.")
		}

		result = append(result, pullrequests...)

		if len(pullrequests) < 100 {
			break
		}
	}
	return result, nil
}

// GetPullRequests returns all pull requests of the repo
func (c *GithubTracker) GetIssues(owner, repo string) ([]*github.Issue, error) {
	result := make([]*github.Issue, 0)
	for page := 0; ; page++ {
		issues, _, err := c.client.Issues.ListByRepo(context.Background(), owner, repo,
			&github.IssueListByRepoOptions{State: "closed", Sort: "created", ListOptions: github.ListOptions{Page: page, PerPage: 100}})
		if err != nil {
			return nil, errors.Wrap(err, "Failed to get issues.")
		}

		result = append(result, issues...)

		if len(issues) < 100 {
			break
		}
	}
	return result, nil
}


