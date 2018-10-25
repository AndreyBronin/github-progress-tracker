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

/*
Package main provides CLI interface based on Cobra and Viper
*/
package main

import (
	log "github.com/sirupsen/logrus"
)

/*
commands:

github-progress-tracker owner add name
github-progress-tracker owner list  - all tracked owners
github-progress-tracker owner repos - all repos of the owner

github-progress-tracker repo list - all tracked repos
github-progress-tracker repo update

options -v  verbose debug log
--storage default sqlite3
 */

func main() {
	err := rootCmd.Execute()
	if err != nil {
		log.Fatalln(err.Error())
	}
}
