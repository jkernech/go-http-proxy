
build:
	@dep ensure
	@go build ./...

run:
	@go run main.go

release:
	@goreleaser --rm-dist

test:
	@dep ensure
	@go test -coverprofile=coverage.out -cover ./...

sonar:
	wget https://sonarsource.bintray.com/Distribution/sonar-scanner-cli/sonar-scanner-2.8.zip
	unzip sonar-scanner-2.8.zip
	@sonar-scanner-2.8/bin/sonar-scanner -Dsonar.organization=$(SONAR_ORGANIZATION) -Dsonar.host.url=$(SONAR_CLOUD_URL) -Dsonar.login=$(SONAR_TOKEN) -e -Dsonar.analysis.mode=publish
	rm -rf sonar-scanner-2.8*
