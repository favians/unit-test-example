export ENV=test
go test -v -coverpkg=./... -coverprofile=coverage.out ./... && go tool cover -func coverage.out && go tool cover -html=coverage.out -o cover.html