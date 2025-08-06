
BINARY_NAME=app

# Makefile contains list of following commands:

# Development

# run tests
run_tests:
	go test ./... -count 1

# formatting
gofmt:
	go fmt ./...

# goimports groups and sorts import sections
goimports:
	goimports --local event-calendar/ -l -w .

#golangci_lint_install:
#	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
#golangci-lint:
#	golangci-lint run

# Generate mocks using mockery
# Prerequisites: mockery.
# Check the latest tool's version on releases page.
mockery_install:
	go install github.com/vektra/mockery/v3@v3.3.2

#generate_mock:
	## example using mockery v2: mockery --name=UserService --dir=internal/service --output=internal/mocks
	#mockery --name=$(INTERFACE_NAME) --dir=$(SOURCE_DIR) --output=$(OUTPUT_DIR)

# Migration tool
# Look at ./cmd/migrator/Makefile to use ready-to-use commands for running database migrations
# Prerequisites: goose.
goose_install:
	go install github.com/pressly/goose/v3/cmd/goose@latest

#----------------------------------------------------------------------------------------------------------

# docker-compose commands for local development
#
# Prerequisites: Docker, docker-compose.

docker_compose_build_and_run:
	docker-compose -f ./docker/docker-compose.yml build && docker-compose -f ./docker/docker-compose.yml up -d

docker_compose_build:
	docker-compose -f ./docker/docker-compose.yml build

docker_compose_run:
	docker-compose -f ./docker/docker-compose.yml up -d

