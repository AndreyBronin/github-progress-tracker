
package tracker

import (
	"github.com/stretchr/testify/assert"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
	"testing"
)

func TestSqliteStorage_SaveCommit(t *testing.T) {
	s, err := NewStorage()
	assert.NoError(t, err)

	err = s.SaveCommit("tmp", &object.Commit{Hash: plumbing.NewHash("70149bd851a131c9afd98ed2a90d3d75068f1218"),
	Message: "initial commit"})
	assert.NoError(t, err)
}