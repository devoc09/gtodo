# Setup scripts
# build binary for Mac and set command into your GOPATH/bin
darwin:
	GOOS=darwin GOARCH=amd64 go build -o bin/mac/gtodo .

# build binary for Linux AMD64 architecture and set command into your GOPATH/bin
linux64:
	GOOS=linux GOARCH=amd64 go build -o bin/linux64/gtodo .

# build binary for Linux ARM architecture and set command into your GOPATH/bin
linuxArm:
	GOOS=linux GOARCH=arm go build -o bin/linuxArm/gtodo .

