// https://cirrus-ci.org/api/
// https://github.com/cirruslabs/cirrus-ci-web/blob/master/schema.graphql

/*

#!/bin/sh

# https://github.com/cirruslabs/cirrus-ci-web/blob/master/schema.graphql

curl -s -X POST --data \
'{
  "query": "query BuildBySHAQuery($owner: String!, $name: String!, $SHA: String) { searchBuilds(repositoryOwner: $owner, repositoryName: $name, SHA: $SHA) { id } }",
  "variables": {
    "owner": "qemu",
    "name": "qemu",
    "SHA": "43d1455cf84283466e5c22a217db5ef4b8197b14"
  }
}' \
https://api.cirrus-ci.com/graphql | python -m json.tool
*/

package backends

import (
	"github.com/ligurio/testres-db/formats"
	"github.com/machinebox/graphql"
	"golang.org/x/net/context"
	"net/http"
)

func SyncCirrusCI(client *http.Client, b *Backend) (*[]formats.TestResult, error) {
	graphql_scheme := "https://api.cirrus-ci.com/graphql"
	ClientOption := graphql.WithHTTPClient(client)
	connection := graphql.NewClient(graphql_scheme, ClientOption)
	request := ""
	req := graphql.NewRequest(request)

	type response struct {
		Name  string
		Items struct {
			Records []struct {
				Title string
			}
		}
	}

	var respData response
	ctx := context.Background()
	if err := connection.Run(ctx, req, &respData); err != nil {
		return nil, err
	}

	return nil, nil
}

/*
type Root {
  viewer: User
  repository(id: ID!): Repository
  githubRepository(owner: String!, name: String!): Repository
  githubRepositories(owner: String!): [Repository]
  githubOrganizationInfo(organization: String!): GitHubOrganizationInfo
  build(id: ID!): Build
  searchBuilds(repositoryOwner: String!, repositoryName: String!, SHA: String): [Build]
  task(id: ID!): Task
  webhookDelivery(id: String!): WebHookDelivery
}

type Build {
  id: ID!
  repositoryId: ID!
  branch: String!
  changeIdInRepo: String!
  changeMessageTitle: String
  changeMessage: String
  durationInSeconds: Int
  clockDurationInSeconds: Int
  pullRequest: Int
  checkSuiteId: Int
  isSenderUserCollaborator: Boolean
  senderUserPermissions: String
  changeTimestamp: Int!
  buildCreatedTimestamp: Int!
  status: BuildStatus
  notifications: [Notification]
  tasks: [Task]
  taskGroupsAmount: Int
  latestGroupTasks: [Task]
  repository: Repository!
  viewerPermission: PermissionType!
}

type Task {
  id: ID!
  buildId: ID!
  repositoryId: ID!
  name: String!
  status: TaskStatus
  notifications: [Notification]
  commands: [TaskCommand]
  artifacts: [Artifacts]
  commandLogsTail(name: String!): [String]
  statusTimestamp: Int!
  creationTimestamp: Int!
  scheduledTimestamp: Int!
  executingTimestamp: Int!
  finalStatusTimestamp: Int!
  durationInSeconds: Int!
  labels: [String]
  uniqueLabels: [String]
  requiredPRLabels: [String]
  optional: Boolean
  statusDurations: [TaskStatusDuration]
  repository: Repository!
  build: Build!
  previousRuns: [Task]
  allOtherRuns: [Task]
  dependencies: [Task]
  automaticReRun: Boolean!
  useComputeCredits: Boolean!
  usedComputeCredits: Boolean!
  transaction: AccountTransaction
  triggerType: TaskTriggerType!
  instanceResources: InstanceResources
}



*/
