APP_NAME		=	Dip
MAJOR_VERSION	=	0
MINOR_VERSION	=	0
BUILD_VERSION	=	1
GIT_COMMIT		=	$(shell git rev-parse --short HEAD)

RELEASE_DIR		=	release
ARTIFACT_NAME	=	$(shell echo $(APP_NAME) | tr A-Z a-z)

IMPORT_VARS		=	-X github.com/feo0o/dip/app.Name=$(APP_NAME) \
					-X github.com/feo0o/dip/app.majorVer=$(MAJOR_VERSION) \
					-X github.com/feo0o/dip/app.minorVer=$(MINOR_VERSION) \
					-X github.com/feo0o/dip/app.buildVer=$(BUILD_VERSION) \
					-X github.com/feo0o/dip/app.gitCommit=$(GIT_COMMIT)

ENV_WINDOWS_X64	=	GOOS=windows GOARCH=amd64
ENV_LINUX_X64	=	GOOS=linux GOARCH=amd64
ENV_DARWIN_X64	=	GOOS=darwin GOARCH=amd64

BUILD_RELEASE	=	go build -trimpath \
					-gcflags="all=-trimpath=$(PWD)" \
					-asmflags="all=-trimpath=$(PWD)" \
					-ldflags '-extldflags "-static" $(IMPORT_VARS)'

release:main.go
	$(ENV_WINDOWS_X64) $(BUILD_RELEASE) -o $(RELEASE_DIR)/$(ARTIFACT_NAME)_windows_amd64.exe main.go
	$(ENV_LINUX_X64) $(BUILD_RELEASE) -o $(RELEASE_DIR)/$(ARTIFACT_NAME)_linux_amd64 main.go
	# $(ENV_DARWIN_X64) $(BUILD_RELEASE) -o $(RELEASE_DIR)/$(ARTIFACT_NAME)_darwin_amd64 main.go

.PHONY:release