    # Go parameters
    GOCMD=go
    GOBUILD=$(GOCMD) build
    GOCLEAN=$(GOCMD) clean
    GOTEST=$(GOCMD) test
    GOGET=$(GOCMD) get
    BINARY_NAME=go-playground
    BINARY_UNIX=$(BINARY_NAME)_unix

    all: build
    build:
	$(GOBUILD) -o ./bin/$(BINARY_NAME) -v ./cmd/...
    #test:
	#$(GOTEST) -v ./...
    clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)
    run:
	$(GOBUILD) -o ./bin/$(BINARY_NAME) -v ./cmd/...
	./bin/$(BINARY_NAME)
	deps:
	$(GOGET) github.com/spf13/viper
	$(GOGET) github.com/gorilla/mux
	$(GOGET) github.com/jinzhu/gorm
	$(GOGET) github.com/jinzhu/gorm/dialects/mysql


# Cross compilation
build-linux:
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v