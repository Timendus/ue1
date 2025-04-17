build: src/*
	@GOOS=linux GOARCH=amd64 go build -o ../dist/linux/ue1asm src/*.go
	@GOOS=windows GOARCH=amd64 go build -o ../dist/windows/ue1asm.exe src/*.go
	@GOOS=darwin GOARCH=amd64 go build -o ../dist/macos-intel/ue1asm src/*.go
	@GOOS=darwin GOARCH=arm64 go build -o ../dist/macos-apple-silicon/ue1asm src/*.go

tests: build
	@go run tests.go

run:
	@go run src/*.go ../programs/fibonacci/input.asm ../programs/fibonacci/output.bin
