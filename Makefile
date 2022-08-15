swagger:
	swag init --dir ./api -g controllers.go --parseDependency

test:
	go test ./... --cover

coverage:
	go test ./... -coverprofile=coverage.out
	cat coverage.out | grep -v "//NOCOVER" > coverage.final.out
	go tool cover -html=coverage.final.out