package backends

import (
	"fmt"
	"github.com/ligurio/testres-db/formats"
	gitlab "github.com/xanzy/go-gitlab"
	"log"
	"net/http"
	"path/filepath"
)

func SyncGitLab(client *http.Client, b *Backend) (*[]formats.TestResult, error) {
	if b.Pipeline != "" {
		log.Println("Option pipeline is specified, but unused")
	}

	gl := gitlab.NewClient(client, b.Secret)
	gl.SetBaseURL(b.Base)

	projOpt := &gitlab.GetProjectOptions{
		Statistics:           gitlab.Bool(false),
		License:              gitlab.Bool(false),
		WithCustomAttributes: gitlab.Bool(false),
	}

	p, _, err := gl.Projects.GetProject(b.Project, projOpt, nil)
	if err != nil {
		return nil, err
	}

	/*

		const (
		    Pending  BuildStateValue = "pending"
		    Running  BuildStateValue = "running"
		    Success  BuildStateValue = "success"
		    Failed   BuildStateValue = "failed"
		    Canceled BuildStateValue = "canceled"
		    Skipped  BuildStateValue = "skipped"
		    Manual   BuildStateValue = "manual"
		)
	*/

	projectOpt := &gitlab.ListProjectPipelinesOptions{
		Scope:   gitlab.String("finished"),
		Status:  gitlab.BuildState(gitlab.Success), // FIXME: use at least failed status
		Ref:     gitlab.String(b.Branch),
		OrderBy: gitlab.String("updated_at"),
		Sort:    gitlab.String("asc"),
	}

	pipelines, _, err := gl.Pipelines.ListProjectPipelines(p.ID, projectOpt)
	if err != nil {
		return nil, err
	}

	jobsOpt := &gitlab.ListJobsOptions{
		ListOptions: gitlab.ListOptions{Page: 1, PerPage: 10},
		Scope:       []gitlab.BuildStateValue{"created", "pending", "running", "failed", "success", "canceled", "skipped"},
	}
	for _, pipeline := range pipelines {
		log.Printf("Found pipeline: %d, status %s", pipeline.ID, pipeline.Status)
		log.Printf("SHA %s, Ref %s", pipeline.SHA, pipeline.Ref)
		jobs, _, err := gl.Jobs.ListPipelineJobs(p.ID, pipeline.ID, jobsOpt, nil)
		if err != nil {
			log.Println(err)
			continue
		}
		for _, job := range jobs {
			log.Println("Found job", job.ID, job.Name, job.Status)
			for _, artifact := range job.Artifacts {
				log.Printf("Found file %s (%d)", artifact.Filename, artifact.Size)
				fnParse := formats.Parser[filepath.Ext(artifact.Filename)]
				if fnParse == nil {
					continue
				}
				fileUrl := artifact.Filename
				if err := DownloadFile(artifact.Filename, fileUrl); err != nil {
					log.Println(err)
					continue
				}
				report, err := fnParse(artifact.Filename)
				if err != nil {
					log.Println(err)
					continue
				}
				report.Name = fmt.Sprintf("%d", job.ID)
				/* FIXME: report.CreatedAt = job.CreatedAt */
			}
		}
	}

	return nil, nil
}
