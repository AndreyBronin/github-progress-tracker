package main

import (
	"github.com/AndreyBronin/github-progress-tracker/tracker"
)


func main() {
	collector, _ := tracker.NewGithubTracker()


	collector.ProcessRepo("bitcoin", "bitcoin")


	/*
	opt := &github.RepositoryListByOrgOptions{Type: "public"}
	repos, _, err := client.Repositories.ListByOrg(context.Background(), "insolar", opt)
	if err != nil {
		log.Fatalln(err.Error())
	}

	for _, r := range repos {
		fmt.Println(*r.Name)
	}

	events, _ , err := client.Activity.ListRepositoryEvents(context.Background(), "insolar", "insolar", &github.ListOptions{0, 500})
	if err != nil {
		log.Fatalln(err.Error())
	}

	for _, e := range events {
		fmt.Println(e.GetCreatedAt(), e.GetType(), string(*e.RawPayload))
	}
*/

	//refs, _, _ := collector.client.Git.ListRefs(context.Background(), "insolar", "insolar", &github.ReferenceListOptions{ListOptions: github.ListOptions{0, 500}})
	//for _, r := range refs {
	//	fmt.Println(r.GetObject().GetType())
	//}
/*
	err := collector.ProcessOrganizationRepos("insolar", []string{"insolar"})
	if err != nil {
		log.Fatalln(err.Error())
	}
*/
	//client.Repositories.ListContributors()
}
