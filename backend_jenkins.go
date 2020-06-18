package backends

import (
	"github.com/bndr/gojenkins"
	"github.com/ligurio/testres-db/formats"
	"log"
	"net/http"
)

func SyncJenkins(client *http.Client, b *Backend) (*[]formats.TestResult, error) {
	var jenkins *gojenkins.Jenkins
	jenkins = gojenkins.CreateJenkins(client, b.Base, b.Username, b.Secret)
	_, err := jenkins.Init()
	if err != nil {
		return nil, err
	}

	jobBuilds, err := jenkins.GetAllBuildIds(b.Pipeline)
	if err != nil {
		return nil, err
	}

	results := make([]formats.TestResult, len(jobBuilds))
	for _, jobBuild := range jobBuilds {
		buildNum, err := jenkins.GetBuild(b.Pipeline, jobBuild.Number)
		if err != nil {
			return &results, err
		}
		log.Println(jobBuild.URL, buildNum.GetResult())
		TestResult, err := buildNum.GetResultSet()
		if err != nil {
			return &results, err
		}
		var testcases []formats.TestCase
		for _, suite := range TestResult.Suites {
			var testcase formats.TestCase
			for _, test := range suite.Cases {
				testcase = formats.TestCase{Name: test.Name}
			}
			testcases = append(testcases, testcase)
		}
		buildInfo := buildNum.Info()
		var result = formats.TestResult{Name: buildInfo.ID, TestCases: testcases}
		results = append(results, result)
	}

	return &results, nil
}
