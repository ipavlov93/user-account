## Back-end service written in GO

https://github.com/ipavlov93/user-account is prototype of back-end service.

### Tests

There are unit and integration tests for this project.
It's recommended to run the tests using this command:
`go test ./... -count 1`

---

## Run prerequisites

#### env file

[//]: # (_.env.example_)
1. Create copy of [.env.local](.env.local) file and put it to _docker/app_ folder.
2. Set values depends on your environment (local, stage, production).

#### Firebase
Create firebase project and setup auth. Export private key _account.json_ file and put it to the project's root.

## docker-compose

#### build and run
`docker-compose -f ./docker/docker-compose.yml up --build -d`

#### run
`docker-compose -f ./docker/docker-compose.yml up -d`

---

### Development

#### formatting

`go fmt ./...`

#### goimports groups import sections

`goimports --local event-calendar/ -l -w .`

#### Linters

`go vet` and golangci-lint is recommended to use.

golangci-lint/v2 bring a new feature to add custom formatters to golangci-lint config.
Formatters will be executed using single command: 
`golangci-lint fmt`.
Formatters are automatically used as “linter” when you run the command `golangci-lint run`.

add [goimports](#goimports-groups-import-sections) as custom formatter to golangci-lint config.

#### Generate mocks using mockery

`mockery --name=UserService --dir=internal/facade --output=internal/mocks`

---

### Project reflection

Application is simple, modular API service designed to register users, provide multiple SSO (Single-Sign-On) methods. 
Built with Go and a clean architecture approach. Application is easy to maintain, extend, and scale.

#### Auth
Firebase auth is used as abstraction for OAuth2.0 (protocol) and to implement (Single-sign-on) authentication flow with multiple auth provider support.

#### Authentication
Application uses Firebase ID tokens for authentication. Tokens should be passed in the Authorization: Bearer <token> header.

Roles and Scopes (planned)
We plan to introduce custom scopes (e.g., users:view, users:create, users: update...) via JWT claims to handle access control across microservices.

### Patterns and architecture style

#### Architecture style

Clean architecture layers:
- domain (reach or anemic models ?)
- repository (dao)
- service
- service facade
- DB adapter
- DTOs

SOLID ?
- DI (technique)
1. DI is used to provide ability to mock any layer(s) in unit tests.
2. dependencies injected via constructor parameters (instead of function call or setter call params). For example, database transaction (abstraction) injected to service/repository.

#### Patterns
- DTOs
- Mapper
- Service facade
