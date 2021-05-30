# github.com/angryronald/guestlist

## Description
Guestlist service for GetGround.com technical test

### Code Style
1. SOLID Principle
2. Dependency Injection
3. Domain Driven Design
4. Command Query Responsibility Segregation

## Test
1. Unit test: /internal/guest/infrastructure/repository/database/guest_test.go
2. Mock test using mockery: /internal/guest/infrastructure/repository/mocks, /internal/guest/domain/service/guest/mocks
3. Integration test: internal/guest/domain/service/guest/guest_test.go

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
