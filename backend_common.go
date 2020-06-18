package backends

import (
	"crypto/tls"
	"errors"
	"github.com/ligurio/testres-db/formats"
	"io"
	"log"
	"net/http"
	"os"
)

type fnSyncBackend func(client *http.Client, b *Backend) (*[]formats.TestResult, error)

var backend = map[string]fnSyncBackend{
	"azure_devops": SyncAzureDevOps,
	"circleci":     SyncCircleCI,
	"cirrusci":     SyncCirrusCI,
	"gitlab":       SyncGitLab,
	"jenkins":      SyncJenkins,
	"teamcity":     SyncTeamCity,
	"travisci":     SyncTravisCI,
}

type Backend struct {
	Name      string
	Base      string
	Project   string
	Branch    string
	Username  string
	Secret    string
	Pipeline  string
	Type      string
	Artifacts []Artifact
}

type Artifact struct {
	Path string
}

var (
	errUnknownBackend = errors.New("Unknown backend")
)

// https://stackoverflow.com/questions/38822764/how-to-send-a-https-request-with-a-certificate-golang/38825553#38825553
func NewAPIClient() *http.Client {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	return client
}

func (b *Backend) GetTestResults() (*[]formats.TestResult, error) {
	log.Println("Backend:", b.Type)
	fn := backend[b.Type]
	if fn == nil {
		return nil, errUnknownBackend
	}
	client := NewAPIClient()
	result, err := fn(client, b)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func DownloadFile(filename string, url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}
