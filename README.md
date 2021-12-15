# spotifas

## Run

### Without Docker

- to run the app, run `nodemon`
- the app should run on port 8080 or the one specified in the PORT environment variable

### With Docker

- if you haven't run the app yet, you need to build the Docker image first by running `docker build .`
- now you can run `docker run` followed by the image id

## Format

To format the code, run `gofmt -w .`

## Lint

The linter used in this app is golangci-lint

To lint the code, run `bin/golangci-lint run`
