.PHONY: build
build:
	go build -o bin/bdn9 main.go
	go build -o bin/bdn9_tray cmd/tray/main.go

.PHONY: proto
proto:
	protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		proto/bdn9.proto

.PHONY: macos
macos:
	iconutil -c icns -o macos/BDN9.app/Contents/Resources/icon.icns icon/bdn9.iconset
	go build -o macos/BDN9.app/Contents/MacOS/bdn9_tray cmd/tray/main.go
