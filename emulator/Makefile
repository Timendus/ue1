build: src/*
	@GOOS=linux GOARCH=amd64 go build -o ../dist/linux/ue1emu src/*.go
	@GOOS=windows GOARCH=amd64 go build -o ../dist/windows/ue1emu.exe src/*.go
	@GOOS=darwin GOARCH=amd64 go build -o ../dist/macos-intel/ue1emu src/*.go
	@GOOS=darwin GOARCH=arm64 go build -o ../dist/macos-apple-silicon/ue1emu src/*.go

run:
	@go run src/*.go ../programs/fibonacci/output.bin
