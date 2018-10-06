
package tracker

import (
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
)
// Storage interface provide methods to store and analyze git commits
type Storage interface {
	SaveCommit(repoName string, commit *object.Commit) error
}

func NewStorage() (Storage, error) {

	db, err := gorm.Open("sqlite3", "commits.db")
	if err != nil {
		return nil, errors.New("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Commit{})


	return &sqliteStorage{db: db, database: "commits.db"}, nil
}


type sqliteStorage struct {
	database string
	db *gorm.DB
}

// Commit struct represents commit record in Database
type Commit struct {
	gorm.Model
	Datetime string
	Name string
	Hash string
}

// SaveCommit inserts commit to database
func (s *sqliteStorage) SaveCommit(repoName string, commit *object.Commit) error {
	fmt.Println(commit.Author.When, commit.Author.Name, commit.Hash)

	s.db.Create(&Commit{Datetime: commit.Author.When.String(),
	Name: commit.Author.Name,
	Hash: commit.Hash.String()})

	return nil
}