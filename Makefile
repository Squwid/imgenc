compile:
	GOOS=darwin GOARCH=amd64 go build -o bin/imgenc_darwin_amd64
	GOOS=darwin GOARCH=arm64 go build -o bin/imgenc_darwin_arm64
	GOOS=windows GOARCH=amd64 go build -o bin/imgenc_windows_amd64
	GOOS=linux GOARCH=amd64 go build -o bin/imgenc_linux_amd64
	GOOS=linux GOARCH=arm64 go build -o bin/imgenc_linux_arm64

clean:
	rm -rf bin/

.PHONY: clean compile