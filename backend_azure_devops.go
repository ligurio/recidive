package backends

import (
	"context"
	"github.com/ligurio/testres-db/formats"
	"github.com/microsoft/azure-devops-go-api/azuredevops"
	"github.com/microsoft/azure-devops-go-api/azuredevops/build"
	"github.com/microsoft/azure-devops-go-api/azuredevops/core"
	"github.com/microsoft/azure-devops-go-api/azuredevops/pipelines"
	"github.com/microsoft/azure-devops-go-api/azuredevops/testresults"
	"log"
	"net/http"
)

func getBuilds(ctx context.Context, connection *azuredevops.Connection, ProjectName *string, BranchName *string) (*[]build.Build, error) {
	buildClient, err := build.NewClient(ctx, connection)
	if err != nil {
		return nil, err
	}

	buildsArgs := build.GetBuildsArgs{Project: ProjectName}
	if *BranchName != "" {
		buildsArgs.BranchName = BranchName
	}
	responseValue, err := buildClient.GetBuilds(ctx, buildsArgs)
	if err != nil {
		return nil, err
	}

	var builds *[]build.Build = nil
	builds = &(*responseValue).Value
	for responseValue != nil {
		// FIXME: builds = append(*builds, &(*responseValue).Value)
		for _, teamBuildReference := range (*responseValue).Value {
			log.Printf("Build %v, %s", *teamBuildReference.BuildNumber, *teamBuildReference.SourceBranch)
		}

		if responseValue.ContinuationToken != "" {
			buildsArgs := build.GetBuildsArgs{
				ContinuationToken: &responseValue.ContinuationToken,
			}
			buildsArgs.ContinuationToken = &responseValue.ContinuationToken
			responseValue, err = buildClient.GetBuilds(ctx, buildsArgs)
			if err != nil {
				return nil, err
			}
		} else {
			responseValue = nil
		}
	}

	testresultsClient := testresults.NewClient(ctx, connection)
	for _, bld := range *builds {
		log.Println("Build Num", *bld.Id)
		TestResultDetailsForBuildArgs := testresults.GetTestResultDetailsForBuildArgs{Project: ProjectName, BuildId: bld.Id}
		TestResults, err := testresultsClient.GetTestResultDetailsForBuild(ctx, TestResultDetailsForBuildArgs)
		if err != nil {
			log.Println("Unsupported?", err)
			continue
		}
		if TestResults != nil {
			log.Println("TestResults", (*TestResults).GroupByField)
			// log.Println("TestResults #%v", (*TestResults).ResultsForGroup)
		}
	}

	return builds, err
}

func getPipelineRef(ctx context.Context, connection *azuredevops.Connection, Project *string, PipelineName *string) (*pipelines.Pipeline, error) {
	pipelineClient := pipelines.NewClient(ctx, connection)
	responseValue, err := pipelineClient.ListPipelines(ctx, pipelines.ListPipelinesArgs{Project: Project})
	if err != nil {
		return nil, err
	}

	var PipelineRef *pipelines.Pipeline = nil
	for _, teamPipelineReference := range (*responseValue).Value {
		// log.Printf("Pipeline = %v", *teamPipelineReference.Name)
		if *teamPipelineReference.Name == *PipelineName {
			PipelineRef = &teamPipelineReference
			break
		}
	}

	return PipelineRef, nil
}

func getProjectRef(ctx context.Context, connection *azuredevops.Connection, ProjectName *string) (*core.TeamProjectReference, error) {
	coreClient, err := core.NewClient(ctx, connection)
	if err != nil {
		return nil, err
	}

	responseValue, err := coreClient.GetProjects(ctx, core.GetProjectsArgs{})
	if err != nil {
		return nil, err
	}

	index := 0
	var Project *core.TeamProjectReference = nil
	for responseValue != nil {
		for _, teamProjectReference := range (*responseValue).Value {
			// log.Printf("Name[%v] = %v", index, *teamProjectReference.Name)
			if *teamProjectReference.Name == *ProjectName {
				Project = &teamProjectReference
				break
			}
			index++
		}

		if responseValue.ContinuationToken != "" {
			projectArgs := core.GetProjectsArgs{
				ContinuationToken: &responseValue.ContinuationToken,
			}
			responseValue, err = coreClient.GetProjects(ctx, projectArgs)
			if err != nil {
				return nil, err
			}
		} else {
			responseValue = nil
		}
	}

	return Project, nil

}

// Using custom http client: https://github.com/microsoft/azure-devops-go-api/issues/52
func SyncAzureDevOps(client *http.Client, b *Backend) (*[]formats.TestResult, error) {
	if b.Username != "" {
		log.Println("Username is specified but unused", b.Username)
	}

	connection := azuredevops.NewPatConnection(b.Base, b.Secret)
	ctx := context.Background()

	project, err := getProjectRef(ctx, connection, &b.Project)
	if err != nil {
		return nil, err
	}
	if project.Url != nil {
		log.Println("URL:", *project.Url)
	}
	log.Println("Last Update Time:", project.LastUpdateTime)

	if project.Abbreviation != nil {
		log.Println(*project.Abbreviation)
	}

	pipeline, err := getPipelineRef(ctx, connection, &b.Project, &b.Pipeline)
	if err != nil {
		return nil, err
	}
	log.Println("Pipeline:", *pipeline.Name)
	builds, err := getBuilds(ctx, connection, project.Name, &b.Branch)
	if err != nil {
		return nil, err
	}

	if builds == nil {
		log.Println("list of builds is empty")
		return nil, err
	}
	for _, bld := range *builds {
		log.Println("Build", *bld.BuildNumber, *bld.SourceBranch)
	}

	return nil, nil
}

/*
http://localhost:6060/pkg/github.com/microsoft/azure-devops-go-api/azuredevops/testplan/
http://localhost:6060/pkg/github.com/microsoft/azure-devops-go-api/azuredevops/test/
http://localhost:6060/pkg/github.com/microsoft/azure-devops-go-api/azuredevops/testresults/
func AzureDevopsTMS_fn(b *Backend) ([]*formats.TestReport, error) {
	log.Println("not implemented")
	return nil, nil
}
*/
