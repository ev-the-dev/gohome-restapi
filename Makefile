GO_FLAGS += "-ldflags=-s -w"
GO_FLAGS += -trimpath

.PHONY build:
build:
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build $(GO_FLAGS)
