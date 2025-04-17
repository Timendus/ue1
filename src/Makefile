build: build-assembler build-emulator

build-assembler: ./*
	@GOOS=linux GOARCH=amd64 go build -o ../dist/linux/ue1asm cmd/assembler/*.go
	@GOOS=windows GOARCH=amd64 go build -o ../dist/windows/ue1asm.exe cmd/assembler/*.go
	@GOOS=darwin GOARCH=amd64 go build -o ../dist/macos-intel/ue1asm cmd/assembler/*.go
	@GOOS=darwin GOARCH=arm64 go build -o ../dist/macos-apple-silicon/ue1asm cmd/assembler/*.go

build-emulator: ./*
	@GOOS=linux GOARCH=amd64 go build -o ../dist/linux/ue1emu cmd/emulator/*.go
	@GOOS=windows GOARCH=amd64 go build -o ../dist/windows/ue1emu.exe cmd/emulator/*.go
	@GOOS=darwin GOARCH=amd64 go build -o ../dist/macos-intel/ue1emu cmd/emulator/*.go
	@GOOS=darwin GOARCH=arm64 go build -o ../dist/macos-apple-silicon/ue1emu cmd/emulator/*.go

tests: build-assembler
	@go run cmd/tests/main.go
