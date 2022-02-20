.PHONY: proto
proto:
	protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		proto/bdn9.proto

macos:
	iconutil -c icns -o macos/BDN9.app/Contents/Resources/icon.icns icon/bdn9.iconset
	go build -o macos/BDN9.app/Contents/MacOS/bdn9_tray cmd/tray/main.go
