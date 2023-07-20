# Bookings & Reservations

Startup:
go run cmd/web/*.go

This is the repository for my bookings & reservations project.

- Built in Go version 1.17
- Uses [chi router](https://github.com/go-chi/chi)
- Uses [edwards](https://github.com/alexedwards/scs)
- Uses [nosurf](https://github.com/justinas/nosurf)

## run local on mac
./run.sh

## run local on windows
run.bat

## run go test in specific folder
go test -v

## run go test coverage in specific folder
go test -cover

## run go test in specific folder with coverage.out converted to html results
go test -coverprofile=coverage.out && go tool cover -html=coverage.out

## run go tests in specific folders that have go files & generate rdm_test_coverage.cov results
/bin/bash -c 'go test -covermode=count -coverprofile=rdm_test_coverage.cov $(go list ./... | grep -v /vendor/)'

## run go tests in specific folders that have go files & generate rdm_test_coverage.cov results html format
/bin/bash -c 'go test -covermode=count -coverprofile=rdm_test_coverage.cov $(go list ./... | grep -v /vendor/) && go tool cover -html=rdm_test_coverage.cov'

### setup tables ###
CREATE DATABASE bookings

CREATE TABLE reservations (id MEDIUMINT NOT NULL AUTO_INCREMENT,room_id int,last_name varchar(255),first_name varchar(255),email varchar(255),phone varchar(255), PRIMARY KEY (id));

CREATE TABLE rooms (id MEDIUMINT NOT NULL AUTO_INCREMENT,name varchar(255),description varchar(255), PRIMARY KEY (id));

CREATE TABLE room_schedules (id MEDIUMINT NOT NULL AUTO_INCREMENT,room_id int, reservation_id int, start_date DATE, end_date DATE, PRIMARY KEY (id));