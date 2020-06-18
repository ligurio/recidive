// https://confluence.jetbrains.com/display/TCD10/REST+API
// https://www.jetbrains.com/help/teamcity/rest-api.html

package backends

import (
	teamcity "github.com/cvbarros/go-teamcity/teamcity"
	"github.com/ligurio/testres-db/formats"
	"log"
	"net/http"
)

func SyncTeamCity(client *http.Client, b *Backend) (*[]formats.TestResult, error) {
	connection, err := teamcity.NewWithAddress(b.Username, b.Secret, b.Base, client)
	if err != nil {
		return nil, err
	}
	project, _ := connection.Projects.GetByID("TestNG_BuildTestsWithGradle")
	if err != nil {
		return nil, err
	}
	log.Println(project)

	// https://godoc.org/github.com/abourget/teamcity#TestOccurrence

	return nil, nil
}
