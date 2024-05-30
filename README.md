# OpsAssist

The project is intended to simplify the life of support engineers in checking configuration files for updates

## Development

### Dependencies

- [Install Go](https://go.dev/doc/install)
- Gitlab API

### Getting started

- Clone this is project on pc
- Go to the project's working directory and download dependencies
  ```bash
  go mod download
  ```
- Make improvements
- Run building
  ```bash
  go build -o opsassist main.go
  ```

## Tests

Full test run
  ```bash
  go test ./tests -v
  ```

Example of running a separate test by name DiffCmd
  ```bash
  go test ./tests -v -run DiffCmd
  ```

## Authors and acknowledgment
Author: [Dmitry Sidyuk](violinorg@yandex.ru)

## [License MIT](./LICENSE)
