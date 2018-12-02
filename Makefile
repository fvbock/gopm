.PHONY: dep build restart run
.EXPORT_ALL_VARIABLES:

BUILD_TARGET=gvault
dep:
	glide install

build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o $(BUILD_TARGET)

run:
	go run main.go

# . ./dev-env && fresh go run main.go
