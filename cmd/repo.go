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

package main

import (
	"github.com/AndreyBronin/github-progress-tracker/tracker"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// repoCmd represents the repo command
var repoCmd = &cobra.Command{
	Use:   "repo",
	Short: "Repo management",

	Run: func(cmd *cobra.Command, args []string) {

		t, err := tracker.NewGithubTracker()
		if err != nil {
			logrus.Fatalln("failed to init tracker")
		}
		repo, err := t.CloneRepo("insolar", "insolar")
		if err != nil {
			logrus.Fatalln("failed to clone repo")
		}


		t.ProcessRepo("insolar", "insolar", repo)
	},
}
