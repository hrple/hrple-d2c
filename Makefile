.PHONY: code-check setup-build-env test only-test only-code-check

build: test
	# ./scripts/build.sh
	echo build

only-build: onlyTest build

package: build package-docker package-zip
	echo "package"

package-docker: 
	# ./scripts/package-docker.sh
	echo "docker"

package-zip: 
	# ./scripts/package-zip.sh
	echo "zip"

code-check: setup-build-env
	# ./scripts/static-code-analysis.sh

only-code-check:
	./scripts/static-code-analysis.sh

only-test:
	rm -f bin/code-coverage-report.html
	go test -v -race -coverprofile=bin/code-coverage-report-company.out -covermode=atomic ./company && go tool cover -html=bin/code-coverage-report-company.out -o bin/code-coverage-report-company.html

test: code-check
	# rm -f bin/code-coverage-report.html
	# go test -v -race -coverprofile=bin/code-coverage-report-company.out -covermode=atomic ./server && go tool cover -html=bin/code-coverage-report-company.out -o bin/code-coverage-report-company.html

setup-build-env:
	# ./scripts/setup-build-env.sh

init:
	./scripts/init.sh