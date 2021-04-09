compile:
	GOOS=darwin GOARCH=amd64 go build -o bin/imgenc_darwin_amd64
	GOOS=darwin GOARCH=arm64 go build -o bin/imgenc_darwin_arm64
	GOOS=windows GOARCH=amd64 go build -o bin/imgenc_windows_amd64
	GOOS=linux GOARCH=amd64 go build -o bin/imgenc_linux_amd64
	GOOS=linux GOARCH=arm64 go build -o bin/imgenc_linux_arm64

zip:
	cd bin && tar -czvf imgenc_darwin_amd64.tar.gz imgenc_darwin_amd64
	cd bin && tar -czvf imgenc_darwin_arm64.tar.gz imgenc_darwin_arm64
	cd bin && tar -czvf imgenc_windows_amd64.tar.gz imgenc_windows_amd64
	cd bin && tar -czvf imgenc_linux_amd64.tar.gz imgenc_linux_amd64
	cd bin && tar -czvf imgenc_linux_arm64.tar.gz imgenc_linux_arm64

mod:
	go mod vendor

clean:
	rm -rf bin/

build: clean mod compile

.PHONY: clean compile mod build