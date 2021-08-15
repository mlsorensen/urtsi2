
default:
	GOOS=linux GOARCH=arm GOARM=7 go build -o urtsi2 main.go

test:
	go test ./...
