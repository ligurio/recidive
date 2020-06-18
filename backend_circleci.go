package backends

// Documentation: https://circleci.com/docs/2.0/artifacts/

import (
	"github.com/jszwedko/go-circleci"
	"github.com/ligurio/testres-db/formats"
	"log"
	"net/http"
	"strings"
)

func SyncCircleCI(client *http.Client, b *Backend) (*[]formats.TestResult, error) {
	project_path := strings.Split(b.Project, "/")
	if len(project_path) != 2 {
		log.Println("Perhaps wrong project name specified")
	}

	account := project_path[0]
	repo := project_path[1]

	connection := &circleci.Client{Token: b.Secret, HTTPClient: client, Debug: true}
	builds, err := connection.ListRecentBuildsForProject(account, repo, b.Branch, "", -1, 0)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	for _, build := range builds {
		log.Printf("Found build: %d, status %s\n", build.BuildNum, build.Status)
		metadata, err := connection.ListTestMetadata(account, repo, build.BuildNum)
		if err != nil {
			log.Println(err)
			return nil, err
		}

		artifacts, err := connection.ListBuildArtifacts(account, repo, build.BuildNum)
		for _, artifact := range artifacts {
			log.Println("Found artifact:", artifact.URL)
		}

		for _, test := range metadata {
			log.Println("Found test:", test.Result, test.Name, test.RunTime)
		}
	}

	return nil, nil
}
