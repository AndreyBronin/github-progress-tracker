/*
 *    Copyright 2018 Andrey Bronin <jonnib@yandex.ru>.
 *
 *    Licensed under the Apache License, Version 2.0 (the "License");
 *    you may not use this file except in compliance with the License.
 *    You may obtain a copy of the License at
 *w
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 *    Unless required by applicable law or agreed to in writing, software
 *    distributed under the License is distributed on an "AS IS" BASIS,
 *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *    See the License for the specific language governing permissions and
 *    limitations under the License.
 */

package tracker

// Score is
type Score uint32

type Counters struct {
	Commits      uint32
	PullRequests uint32
	Issues uint32
}

type Coefficients struct {
	Commits      uint32
	PullRequests uint32
	Issues uint32
}

// CalculateScore calculates rating score for period
// simple formula without coefficients
func CalculateScore(counters Counters, cf Coefficients) uint32 {
	return counters.Commits * cf.Commits +
		counters.PullRequests  * cf.PullRequests +
		counters.Issues  * cf.Issues
}
