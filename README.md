# Bookings & Reservations

Startup:
go run cmd/web/*.go

This is the repository for my bookings & reservations project.

- Built in Go version 1.17
- Uses [chi router](https://github.com/go-chi/chi)
- Uses [edwards](https://github.com/alexedwards/scs)
- Uses [nosurf](https://github.com/justinas/nosurf)

## run go test in specific folder
go test -v

## run go test in specific folder with coverage.out converted to html results
go test -coverprofile=coverage.out && go tool cover -html=coverage.out

## run go tests in specific folders that have go files & generate rdm_test_coverage.cov results
/bin/bash -c 'go test -covermode=count -coverprofile=rdm_test_coverage.cov $(go list ./... | grep -v /vendor/)'

## run go tests in specific folders that have go files & generate rdm_test_coverage.cov results html format
/bin/bash -c 'go test -covermode=count -coverprofile=rdm_test_coverage.cov $(go list ./... | grep -v /vendor/) && go tool cover -html=rdm_test_coverage.cov'