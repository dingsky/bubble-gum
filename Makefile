
OS=""
ifeq ($(OS),"")

else 
	GOCOMPILE = GOOS=$(OS)
endif
mock:
	$(GOCOMPILE) go build -o bin/channelMock channelMock/cmd/main.go
