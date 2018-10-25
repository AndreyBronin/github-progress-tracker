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
	"github.com/stretchr/testify/assert"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
	"testing"
	"time"
)

func TestSqliteStorage_SaveCommit(t *testing.T) {
	t.Skip("tmp")
	s, err := NewStorage("~/commits_.db")
	assert.NoError(t, err)

	err = s.SaveCommit("tmp", &object.Commit{
		Hash:    plumbing.NewHash("70149bd851a131c9afd98ed2a90d3d75068f1218"),
		Message: "initial commit",
		Author:  object.Signature{Name: "Andrey", Email: "info@example.com", When: time.Now()},
	})
	assert.NoError(t, err)


}

