package backends

import (
	"context"
	"fmt"
	"github.com/ligurio/testres-db/formats"
	travisci "github.com/shuheiktgw/go-travis"
	"log"
	"net/http"
	"path"
)

func SyncTravisCI(client *http.Client, b *Backend) (*[]formats.TestResult, error) {
	connection := travisci.NewClient(b.Base, b.Secret)
	connection.HTTPClient = client
	build_service := connection.Builds

	ctx := context.Background()
	var options travisci.BuildsOption
	builds, _, err := build_service.List(ctx, &options)
	if err != nil {
		return nil, err
	}

	results := make([]formats.TestResult, len(builds))
	baseURL := "https://travis-ci.org/"
	for _, build := range builds {
		// https://godoc.org/github.com/shuheiktgw/go-travis#Build
		metadata := *build.Metadata
		log.Printf("Found build: %s, status %s\n", path.Join(baseURL, b.Pipeline, *metadata.Href), *build.State)
		buildId := fmt.Sprintf("%d", build.Id)
		var testcases []formats.TestCase
		var result = formats.TestResult{Name: buildId}
		log.Println("BuildOn", *build.FinishedAt)
		//var result = formats.TestResult{Name: buildId, CreatedAt: *build.FinishedAt}
		var testcase formats.TestCase
		for _, job := range build.Jobs {
			// https://godoc.org/github.com/shuheiktgw/go-travis#Job
			if job.State == nil {
				continue
			}
			log.Println("Job ID: ", *job.Id)
			testcase = formats.TestCase{Name: "XXX"}
		}
		testcases = append(testcases, testcase)
		result.TestCases = testcases
		results = append(results, result)
	}

	return nil, nil
}
