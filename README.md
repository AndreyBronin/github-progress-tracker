[![Build Status](https://travis-ci.org/AndreyBronin/github-progress-tracker.svg?branch=master)](https://travis-ci.org/AndreyBronin/github-progress-tracker)
[![Go Report Card](https://goreportcard.com/badge/github.com/AndreyBronin/github-progress-tracker)](https://goreportcard.com/report/github.com/AndreyBronin/github-progress-tracker)
[![codecov](https://codecov.io/gh/andreybronin/github-progress-tracker/branch/master/graph/badge.svg)](https://codecov.io/gh/andreybronin/github-progress-tracker)
[![GoDoc](https://godoc.org/github.com/andreybronin/github-progress-tracker?status.svg)](https://godoc.org/github.com/andreybronin/github-progress-tracker)


Github progress tracker
=======================

There are a lot of thematic websites witch rates open-source projects. 
Rating usually based on development progress metrics: Github stars, forks, open issues, count of commits and contributors, etc. 

**Github progress tracker** is a tool to do the same thing.
It follow the repo progress and collects development metrics.


#### Workflow

1. Tracker clones selected repo and then stores all commits history to sqlite database cache.
2. Tracker uses github API to get info about: pull requests, issues, releases, contributors, etc.
3. tracker calculates progress score for each development day.
4. static website generator is used to make web-frontend


#### Contributing

You can take any task you like on project [Kanban board](https://github.com/AndreyBronin/github-progress-tracker/projects/2).
Fork the repo, write code and unit tests, create pull.
Don't forget to star the repo. 