## Back-end service written in GO

https://github.com/ipavlov93/user-account is prototype of back-end service.

### Tests

There are unit and integration tests for this project.
It's recommended to run the tests using this command:
`go test ./... -race -count 1`

---

## Run prerequisites

#### env file

[//]: # (_.env.example_)
1. Create copy of [.env.local](.env.local) file and put it to _docker/app_ folder.
2. Set values depends on your environment (local, stage, production).

#### Firebase
Create firebase project and setup auth. Export private key _account.json_ file and put it to the project's root.

## docker-compose

You can find more commands in [Makefile](Makefile).

#### build and run
`
docker-compose -f ./docker/docker-compose.yml build && docker-compose -f ./docker/docker-compose.yml up -d
`

#### run
`
docker-compose -f ./docker/docker-compose.yml up -d
`

---

### Development

#### formatting

`go fmt ./...`

#### goimports groups and sorts import sections

`goimports --local user-account/ -l -w .`

#### Linters

`go vet` and golangci-lint is recommended to use.

Recommendation:
golangci-lint/v2 bring a new feature to add custom formatters to golangci-lint config.
Formatters will be executed using single command: 
`golangci-lint fmt`.
Formatters are automatically used as “linter” when you run the command `golangci-lint run`.

add [goimports](#goimports-groups-import-sections) as custom formatter to golangci-lint config.

#### Generate mocks using mockery

`mockery --name=YourService --dir=internal/your_path --output=internal/mocks`

---

### Migration tool

Prerequisites: [goose](https://github.com/pressly/goose).

Installation:
`
go install github.com/pressly/goose/v3/cmd/goose@latest
`

Look at [Makefile](./cmd/migrator/Makefile) that contains ready-to-use commands for running database migrations.
