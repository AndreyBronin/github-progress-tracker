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

package storage

import (
	"github.com/jinzhu/gorm"
	"time"
)

// Owner struct is gorm model which represents github owner in Database
type Owner struct {
	gorm.Model
	Repositories   []Repository
	Name string
	URL string
}

// Commit struct is gorm model which represents commit record in Database
type Repository struct {
	gorm.Model
	Owner Owner `gorm:"foreignkey:OwnerRefer"`
	Name         string
}

// Score struct is gorm model which represents score record in Database
type Score struct {
	gorm.Model
	Owner Owner `gorm:"foreignkey:OwnerRefer"`
	Repository    Repository `gorm:"foreignkey:RepositoryRefer"`
	Datetime   time.Time
	Score         uint32
	CommitCount uint32
	PullRequestCount uint32
	IssueCount uint32
}

// Commit struct is gorm model which represents commit record in Database
type Commit struct {
	gorm.Model
	Repository Repository `gorm:"foreignkey:RepositoryRefer"`
	Datetime   time.Time
	Name       string
	Hash       string
}

// Contributor struct is gorm model which represents contributor record in Database
type Contributor struct {
	gorm.Model
	Repository    Repository `gorm:"foreignkey:RepositoryRefer"`
	GithubID      uint32
	Login         string
	URL           string
	AvatarURL     string
	Contributions uint32
}
