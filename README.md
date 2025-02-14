# go test commands
go test -coverprofile cover.out
	#Below command to generate htmlfile
go tool cover -html=cover.out -o cover.html
