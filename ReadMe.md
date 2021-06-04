# Guestlist

## Description
Guestlist service

### Code Style
1. SOLID Principle
2. Dependency Injection
3. Domain Driven Design
4. Command Query Responsibility Segregation

## Test
1. Unit test using mockery: /internal/guest/infrastructure/repository/mocks, /internal/guest/domain/service/guest/mocks
2. Integration test: internal/guest/domain/service/guest/guest_test.go, /internal/guest/infrastructure/repository/database/guest_test.go

## Requirement
1. MySQL
3. Go 1.16

## How to run 

### To run the service:
`go run cmd/guestlist/main.go`

### Exposed port:
Listen to port 8001 by default

## Configuration:
Change configuration can be set by set up environment variables or update config/config.go file.

## Build Docker Image
1. Run command docker build --label "version={version}" -t {name}:{version} .
2. Example: docker build --label "version=v1.0.0" -t guestlist:v1.0.0 .