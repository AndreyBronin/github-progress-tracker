/*
 *    Copyright © 2018 Andrey Bronin <jonnib@yandex.ru>.
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

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/pkg/errors"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
)

// Storage interface provide methods to store cache and analyze git commits
type Storage interface {
	SaveCommit(repoName string, commit *object.Commit) error
}

func NewStorage(filePath string) (Storage, error) {

	db, err := gorm.Open("sqlite3", filePath)
	if err != nil {
		return nil, errors.New("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Repository{}, &Commit{}, &Contributor{})
	return &sqliteStorage{db: db, database: filePath}, nil
}

type sqliteStorage struct {
	database string
	db       *gorm.DB
}

// Commit struct is gorm model which represents commit record in Database
type Repository struct {
	gorm.Model
	Name string
}

// Commit struct is gorm model which represents commit record in Database
type Commit struct {
	gorm.Model
	Repository Repository `gorm:"foreignkey:RepositoryRefer"`
	Datetime   string
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

// SaveCommit inserts commit to database
func (s *sqliteStorage) SaveCommit(repoName string, commit *object.Commit) error {
	fmt.Println(commit.Author.When, commit.Author.Name, commit.Hash)

	s.db.Create(&Commit{
		Datetime: commit.Author.When.String(),
		Name:     commit.Author.Name,
		Hash:     commit.Hash.String(),
	})

	return nil
}

// SaveCommit inserts contributor to database
func (s *sqliteStorage) SaveContributor(repoName string, commit *object.Commit) error {
	fmt.Println(commit.Author.When, commit.Author.Name, commit.Hash)

	s.db.Create(&Commit{
		Datetime: commit.Author.When.String(),
		Name:     commit.Author.Name,
		Hash:     commit.Hash.String(),
	})

	return nil
}
