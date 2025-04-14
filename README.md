## Back-end service written in GO

https://github.com/ipavlov93/event-calendar

### Tests

There are unit and integration tests for this project.
It's recommended to run the tests using this command:
`go test ./... -count 1`

---

## Run prerequisites

#### env file

[//]: # (_.env.example_)
1. Create copy of [.env.example](.env.example) file and put it to _docker/_ folder.
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

#### Generate mocks using mockery

`mockery --name=UserService --dir=internal/facade --output=internal/mocks`

---

### Project reflection

#### Auth

Firebase auth is used as abstraction for OAuth2.0 (protocol) and to implement (Single-sign-on) authentication flow with multiple auth provider support.

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
- Mapper ?
- Service facade
