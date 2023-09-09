LIB_FOLDER = ./statscollector
CONFIG_FOLDER = ./config

build: $(LIB_FOLDER)
	CONFIG_FOLDER=$(CONFIG_FOLDER) go build -a -o bin/statscollector cmd/statscollector.go

run: $(LIB_FOLDER)
	CONFIG_FOLDER=$(CONFIG_FOLDER) go run cmd/statscollector.go

test: $(LIB_FOLDER)
	CONFIG_FOLDER=../$(CONFIG_FOLDER) go test -v $(LIB_FOLDER)


