### TODO

- конверторы для форматов в `formats/common.go`
- SubUnit V2
  - https://github.com/msgpack/msgpack-go/blob/master/unpack.go
  - https://github.com/hashicorp/go-msgpack/blob/master/codec/decode.go
- импорт результатов только старше даты последнего результата из базы
- добавлять кастомные сертификаты для http клиента (`ca.go`)
- сохранять в структуре бранч, название пайплайна, коммиты
- поддержка Lava
- поддержка Cirrus CI
- опция `-limit`

### GitHub Actions

- https://help.github.com/en/actions/configuring-and-managing-workflows/persisting-workflow-data-using-artifacts
- https://developer.github.com/v3/actions/
- ? https://github.com/google/go-github

### Zuul

- Example: https://zuul.opendev.org/t/openstack/builds?project=openstack/glance
- Example: https://zuul.opendev.org/t/openstack/builds?project=openstack/ceilometer
- Example: https://zuul.opendev.org/t/openstack/builds?project=openstack/heat
- https://zuul.opendev.org/openapi

### BitBucket

- https://github.com/ktrysmt/go-bitbucket

### Lava

- https://staging.validation.linaro.org/api/help/
- XML RPC Client https://github.com/kolo/xmlrpc

### Kernel CI

- https://api.kernelci.org
- https://github.com/kernelci/kernelci-backend
- Where is a Golang API?
- token is required

### Codefresh

- https://github.com/codefresh-io/go-sdk
- Example: https://github.com/nemequ/portable-snippets
- test results unavailable via API

### Patchwork/Patchew

- https://patchwork-freedesktop.readthedocs.io/en/latest/rest.html
- https://github.com/patchew-project/patchew
- Where is Golang API?

### TestRail

- http://docs.gurock.com/testrail-api2/start
- https://github.com/gurock/testrail-api
- Go: https://github.com/educlos/testrail
- Go: https://godoc.org/github.com/Etienne42/testrail
- https://secure.gurock.com/customers/testrail/trial/
- github.com/educlos/testrail

### Beaker

- https://beaker-project.org/docs/server-api/
- Where is Golang API?

### BuildBot

- REST API: https://github.com/buildbot/buildbot/blob/master/master/docs/developer/rest.rst
- Example: https://github.com/buildbot/buildbot/wiki/SuccessStories
- Example: http://buildbot.suricata-ids.org
- Example: https://buildbot.python.org/
- Example: http://212.201.121.110:38010/
- Example: https://buildbot.openinfosecfoundation.org/
- Example: https://ci.chromium.org/p/chromium/g/main/console
- Example https://chromium.googlesource.com/infra/luci/luci-go/+/master/grpc/prpc/talk/buildbot/client/main.go
- LUCI: http://bit.ly/2kgyE9U
- https://docs.buildbot.net/latest/developer/rest.html
- "go.chromium.org/luci/grpc/prpc/talk/buildbot/proto"
- "go.chromium.org/luci/milo/buildsource/buildbot"
- "go.chromium.org/luci/milo/buildsource/buildbot/buildbotapi"
- "go.chromium.org/luci/milo/buildsource/buildbot/buildstore"

### Drone CI

- Publish test results in artifacts: https://github.com/drone/docs.drone.io/blob/master/artifacts.markdown
- Enterprise only: https://0-8-0.docs.drone.io/publish-unit-test-results/
- API: https://docs.drone.io/api/endpoints/builds/build_list/
- https://github.com/drone/drone/issues/239
- Golang API: https://github.com/drone/drone-go

### CDash

- https://my.cdash.org/viewProjects.php
- https://www.paraview.org/Wiki/CDash:API
- https://open.cdash.org/viewProjects.php
- https://open.cdash.org/viewTest.php?buildid=6227968
- https://open.cdash.org/viewTest.php?buildid=6227571
- https://my.cdash.org/viewTest.php?buildid=1735823
- Where is Golang API?

### AWS CodePipeline

- https://docs.aws.amazon.com/en_us/codepipeline/latest/userguide/welcome.html
- GoDoc: https://docs.aws.amazon.com/sdk-for-go/api/service/codepipeline/
- GoDoc: https://docs.aws.amazon.com/sdk-for-go/api/service/codepipeline/#Artifact

### Appveyor

- API: https://www.appveyor.com/docs/api/projects-builds/
- Can't access to test results via REST API https://github.com/appveyor/ci/issues/3226
- Where is a Golang API? https://github.com/appveyor/ci/issues/3225
- Example: https://ci.appveyor.com/project/rpcs3/rpcs3/branch/master/tests
- Example: https://ci.appveyor.com/project/dignifiedquire/deltachat-core-rust/branch/master/tests
- Example: https://ci.appveyor.com/project/quixdb/portable-snippets/branch/master

### CodeShip CI

- https://apidocs.codeship.com/v2/introduction/basic-vs-pro
- GoDoc: https://godoc.org/github.com/codeship/codeship-go#Build
- How to get a testing results?

### Atlassian Bamboo

- https://developer.atlassian.com/server/bamboo/rest-apis/
- https://github.com/rcarmstrong/go-bamboo
- TODO: How to get test results?
- https://community.atlassian.com/t5/Answers-Developer-Questions/How-to-Get-Bamboo-Test-Results-via-REST/qaq-p/475715

### Concourse CI

- https://ci.spearow.io/teams/main/pipelines/oregano/jobs/make-release/builds/30

### Report Portal

- Where is an API documentation? https://github.com/reportportal/service-api/issues/1094
- Go API https://github.com/avarabyeu/goRP
- Example: https://rp.epam.com/ui/
- Example: http://web.demo.reportportal.io/ui/
- https://github.com/ihar-kahadouski/dev-guide/blob/master/reporting.md

### JetBrains Space

- Golang API?
- Example?

### Bitrise

- https://api-docs.bitrise.io/
- https://devcenter.bitrise.io/testing/test-reports/
- https://devcenter.bitrise.io/testing/exporting-to-test-reports-from-custom-script-steps/
